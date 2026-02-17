package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"email-dispatcher/shared/models"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Failed to Connect to RabbitMQ:", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel", err)
	}
	defer ch.Close()

	//Queue Declare
	q, err := ch.QueueDeclare(
		"email_queue", //name
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to declare queue", err)
	}

	apps := []string{"app-1", "app-2", "app-3"}

	for _, appID := range apps {
		go publishEmails(appID, ch, q.Name)
	}

	//Always open main thread
	select {}
}

func publishEmails(appID string, ch *amqp.Channel, queueName string) {
	for i := 1; i <= 5; i++ {
		email := models.EmailEvent{
			AppId:     appID,
			EmailId:   fmt.Sprintf("%s-email-%d", appID, i),
			Recipient: fmt.Sprintf("user%d@example.com", i),
			Subject:   fmt.Sprintf("Test Email %d from %s", i, appID),
			Body:      fmt.Sprintf("This is test email %d from %s", i, appID),
			TimeStamp: time.Now().Unix(),
		}

		body, _ := json.Marshal(email)

		err := ch.Publish(
			"",        // exchange
			queueName, // routing key
			false,     // mandatory
			false,     // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
				Timestamp:   time.Now(),
			})

		if err != nil {
			log.Printf("[%s] Failed to publish: %v", appID, err)
			continue
		}

		log.Printf("[%s] Email %d sent: %s", appID, i, email.EmailId)
		time.Sleep(1500 * time.Millisecond)
	}
}
