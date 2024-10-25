# goEmail [![GoDoc](https://pkg.go.dev/badge/goEmail.svg)](https://pkg.go.dev/github.com/richi0/goEmail)

goEmail sends emails using different email providers

## What does goEmail do?

goEmail sends emails using different email providers. At the moment two providers are supportet.

- Mailjet: For production use
- Terminal: For development (logs to the terminal)

## How do I use goEmail?

### Install

```
go get -u github.com/richi0/goEmail
```

### Usage Example

```go
// Use different sender depending on your stage
var emailSender goEmail.EmailSender
if stage == "production" {
	emailSender = goEmail.NewMailjetSender("<public_key>", "<private_key>")
} else {
	emailSender = goEmail.NewTerminalSender()
}
emailClient := goEmail.NewEmailClient(emailSender)

err := emailClient.SendEmail(
	"admin@example.com",
	"goEmail",
	"admin@goemail.com",
	"New Alert",
	"Alert message!"
)
```

## Documentation

Find the full documentation of the package here: https://pkg.go.dev/github.com/richi0/goEmail
