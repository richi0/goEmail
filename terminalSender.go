package goEmail

import "fmt"

// A TerminalSender represents a client that can send emails to the terminal.
type TerminalSender struct {
}

// A NewTerminalSender creates a new TerminalSender.
func NewTerminalSender() *TerminalSender {
	return &TerminalSender{}
}

// SendEmail sends an email to the terminal.
func (t *TerminalSender) SendEmail(
	to string,
	fromName string,
	fromEmail string,
	subject string,
	body string,
) error {
	fmt.Printf(
		"Terminal: Sending email to %s from %s <%s> with subject %s and body %s\n",
		to,
		fromName,
		fromEmail,
		subject,
		body,
	)
	return nil
}
