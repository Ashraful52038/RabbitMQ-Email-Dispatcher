package mailpit

import (
	"bytes"
	"email-dispatcher/shared/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Client struct {
	APIEndpoint string
	HTTPClient  *http.Client
}

func NewClient() *Client {
	return &Client{
		APIEndpoint: "http://localhost:8025",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// Mailpit API expects this format
type MailpitAddress struct {
	Email string `json:"Email"`
}

type MailpitMessage struct {
	To      []MailpitAddress `json:"To"`
	From    MailpitAddress   `json:"From"`
	Subject string           `json:"Subject"`
	Text    string           `json:"Text"`
}

func (m *Client) SendEmail(email models.EmailEvent) error {
	log.Printf("[Mailpit] Sending to: %s, Subject: %s", email.Recipient, email.Subject)

	time.Sleep(500 * time.Microsecond)

	message := MailpitMessage{
		To: []MailpitAddress{
			{Email: email.Recipient},
		},
		From: MailpitAddress{
			Email: "sender@example.com", // You can make this dynamic later
		},
		Subject: email.Subject,
		Text:    email.Body,
	}

	jsonDta, _ := json.Marshal(message)
	resp, err := http.Post(m.APIEndpoint+"/api/v1/send", "application/json", bytes.NewBuffer(jsonDta))
	if err != nil {
		log.Printf("[Mailpit] HTTP request error: %v", err)
		return fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Check response
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("mailpit API returned %d: %s", resp.StatusCode, string(body))
	}

	log.Printf("[Mailpit] Successfully sent email to %s", email.Recipient)

	return nil
}
