package mailpit

import (
	"bytes"
	"email-dispatcher/shared/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Client struct {
	APIEndpoint string
}

func NewClient() *Client {
	return &Client{
		APIEndpoint: "http://localhost:8025",
	}
}

func (m *Client) SendEmail(email models.EmailEvent) error {
	log.Printf("[Mailpit] Sending to: %s, Subject: %s", email.Recipient, email.Subject)

	time.Sleep(500 * time.Microsecond)

	emailData := map[string]interface{}{
		"to":      email.Recipient,
		"subject": email.Subject,
		"text":    email.Body,
	}
	jsonDta, _ := json.Marshal(emailData)
	resp, err := http.Post(m.APIEndpoint+"api/v1/send", "application/json", bytes.NewBuffer(jsonDta))
	if err != nil {
		log.Printf("[Mailpit] HTTP request error: %v", err)
		return fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	return nil
}
