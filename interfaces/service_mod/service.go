package service_mod

import "fmt"

// 1. Define an interface contract
type MessageSender interface {
	SendMessage(recipient, body string) error
}

// 2. Concrete implementation (e.g., Production Email Service)
type EmailService struct {
	SMTPServer string
}

func (e EmailService) SendMessage(recipient, body string) error {
	fmt.Printf("Sending Email via %s to %s: %s\n", e.SMTPServer, recipient, body)
	return nil
}

// 3. Consumer depends ONLY on the interface
type NotificationManager struct {
	Sender MessageSender // Decoupled dependency
}

func (nm NotificationManager) NotifyUser(userID, message string) {
	_ = nm.Sender.SendMessage(userID, message)
}
