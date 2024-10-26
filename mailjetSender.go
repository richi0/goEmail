// Package goEmail provides methods to send emails using different email services.
// The two implementations of the EmailSender interface are TerminalSender and MailjetSender.
package goEmail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// A MailjetSender represents a client that can send emails using the Mailjet service.
type MailjetSender struct {
	MjApiKeyPublic  string
	MjApiKeyPrivate string
	client          *http.Client
}

// NewMailjetSender creates a new MailjetSender.
func NewMailjetSender(mjApiKeyPublic string, mjApiKeyPrivate string) *MailjetSender {
	return &MailjetSender{
		MjApiKeyPublic:  mjApiKeyPublic,
		MjApiKeyPrivate: mjApiKeyPrivate,
		client:          &http.Client{},
	}
}

// A Contact represents an email contact.
type Contact struct {
	Email string
	Name  string
}

// A Message represents an email message.
type Message struct {
	From     *Contact
	To       []Contact
	Subject  string
	TextPart string
	HTMLPart string
}

// A MailBody represents the body of an email.
type MailBody struct {
	Messages []Message
}

// SendEmail sends an email to the specified recipient using the Mailjet service.
func (m *MailjetSender) SendEmail(
	to string,
	fromName string,
	fromEmail string,
	subject string,
	body string,
) error {
	var messages []Message
	contact := Contact{Email: to, Name: to}
	message := Message{
		From: &Contact{
			Name:  fromName,
			Email: fromEmail},
		To:       []Contact{contact},
		Subject:  subject,
		TextPart: "",
		HTMLPart: body,
	}
	messages = append(messages, message)
	bodyData, err := json.Marshal(&MailBody{Messages: messages})
	if err != nil {
		return err
	}
	r, err := http.NewRequest("POST", "https://api.mailjet.com/v3.1/send", bytes.NewBuffer(bodyData))
	if err != nil {
		return err
	}
	r.SetBasicAuth(m.MjApiKeyPublic, m.MjApiKeyPrivate)
	res, err := m.client.Do(r)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	if res.StatusCode > 299 {
		body, _ := io.ReadAll(res.Body)
		fmt.Println(res.StatusCode, string(body))
		return err
	}
	return nil
}
