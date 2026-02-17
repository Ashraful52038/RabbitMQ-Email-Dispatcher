package models

type EmailEvent struct {
	AppId     string `json:"app_id"`
	EmailId   string `json:"email_id"`
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	TimeStamp int64  `json:"timestamp"`
}
