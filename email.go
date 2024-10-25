// Package goEmail provides methods to send emails using different email services.
package goEmail

// A EmailClient represents a client that can send emails.
type EmailClient struct {
	sender EmailSender
}

// EmailSender is an interface that defines the method to send emails.
type EmailSender interface {
	SendEmail(
		to string,
		fromName string,
		fromEmail string,
		subject string,
		body string,
	) error
}

// NewEmailClient creates a new EmailClient.
func NewEmailClient(sender EmailSender) *EmailClient {
	return &EmailClient{
		sender: sender,
	}
}

// SendEmail sends an email to the specified recipient.
func (e *EmailClient) SendEmail(
	to string,
	fromName string,
	fromEmail string,
	subject string,
	body string,
) error {
	return e.sender.SendEmail(
		to,
		fromName,
		fromEmail,
		subject,
		body,
	)
}
