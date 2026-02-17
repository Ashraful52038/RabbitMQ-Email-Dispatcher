package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"email-dispatcher/consumer/mailpit"
	"email-dispatcher/shared/models"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5673/")
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open channel:", err)
	}
	defer ch.Close()

	// QoS set
	ch.Qos(1, 0, false)

	// Queue Declare
	q, err := ch.QueueDeclare(
		"email_queue", true, false, false, false, nil,
	)
	if err != nil {
		log.Fatal("Failed to declare queue:", err)
	}

	// Phase 2 worker
	numWorkers := 2
	log.Printf("Starting %d workers for Phase 2", numWorkers)

	mailpitClient := mailpit.NewClient()

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Message consume
	msgs, err := ch.Consume(
		q.Name, "", false, false, false, false, nil,
	)
	if err != nil {
		log.Fatal("Failed to register consumer:", err)
	}

	// Worker pool started
	for i := 0; i < numWorkers; i++ {
		go worker(i+1, &wg, msgs, ch, mailpitClient)
	}

	log.Println("Consumer started. Press Ctrl+C to exit.")

	// Graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	log.Println("Shutting down...")
	ch.Close()
	wg.Wait()
}

func worker(id int, wg *sync.WaitGroup, msgs <-chan amqp.Delivery, ch *amqp.Channel, mailpit *mailpit.Client) {
	defer wg.Done()

	log.Printf("Worker %d started", id)

	for msg := range msgs {
		log.Printf("[Worker %d] Waiting 2 seconds before processing", id)
		time.Sleep(2 * time.Second)

		var email models.EmailEvent
		json.Unmarshal(msg.Body, &email)

		log.Printf("[Worker %d] Processing: %s", id, email.EmailId)

		// send mailpit
		err := mailpit.SendEmail(email)
		if err != nil {
			log.Printf("[Worker %d] Failed: %v", id, err)
			msg.Nack(false, true) // requeue
			continue
		}

		msg.Ack(false)
		log.Printf("[Worker %d] Completed: %s", id, email.EmailId)
	}
}
