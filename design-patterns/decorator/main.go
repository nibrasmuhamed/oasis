// The Decorator Pattern in Go allows you to dynamically add behavior to
// objects without modifying their structure, providing flexibility in
// extending functionality. This is useful for cases like logging,
// authentication, or modifying the behavior of services. In Go,
// functions and interfaces can be decorated using this pattern.

// Key Concepts:

// •	Component Interface: The interface that defines the behavior you want to extend or decorate.
// •	Concrete Component: The basic implementation of the interface.
// •	Decorator: The object that wraps a concrete component and adds additional functionality.
package main

import "fmt"

// Notifier is the component interface
type Notifier interface {
	Send(message string) string
}

// EmailNotifier is the concrete implementation of Notifier
type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) string {
	return fmt.Sprintf("Sending email with message: %s", message)
}

// Decorator adds extra functionality to Notifier
type NotifierDecorator struct {
	Notifier Notifier
}

func (d *NotifierDecorator) Send(message string) string {
	return d.Notifier.Send(message)
}

// SMSDecorator is a concrete decorator that adds SMS functionality
type SMSDecorator struct {
	Notifier Notifier
}

func (s *SMSDecorator) Send(message string) string {
	smsMessage := fmt.Sprintf("Sending SMS with message: %s", message)
	return fmt.Sprintf("%s\n%s", s.Notifier.Send(message), smsMessage)
}

// SlackDecorator is another concrete decorator adding Slack functionality
type SlackDecorator struct {
	Notifier Notifier
}

func (s *SlackDecorator) Send(message string) string {
	slackMessage := fmt.Sprintf("Sending Slack message: %s", message)
	return fmt.Sprintf("%s\n%s", s.Notifier.Send(message), slackMessage)
}

func main() {
	// Basic notifier
	emailNotifier := &EmailNotifier{}

	// Adding SMS functionality
	smsNotifier := &SMSDecorator{Notifier: emailNotifier}

	// Adding Slack functionality
	slackNotifier := &SlackDecorator{Notifier: smsNotifier}

	// Sending messages
	fmt.Println(slackNotifier.Send("Hello World"))
}
