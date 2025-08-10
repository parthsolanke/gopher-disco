package main

import (
	"fmt"
	"time"
)

// Like a personal assistant managing multiple communication channels - email, phone, and instant messages - responding to whichever needs attention first.

func emailManager(emails chan string) {
	emailList := []string{
		"📧 Meeting reminder from boss",
		"📧 Newsletter subscription",
		"📧 Project update from team",
	}

	for _, email := range emailList {
		time.Sleep(2 * time.Second)
		emails <- email
	}
	close(emails)
}

func phoneManager(calls chan string) {
	callList := []string{
		"📞 Call from client about project",
		"📞 Call from dentist office",
	}

	for _, call := range callList {
		time.Sleep(3 * time.Second)
		calls <- call
	}
	close(calls)
}

func instantMessages(messages chan string) {
	msgList := []string{
		"💬 Slack message from colleague",
		"💬 WhatsApp from friend",
		"💬 Teams notification",
	}

	for _, msg := range msgList {
		time.Sleep(1 * time.Second)
		messages <- msg
	}
	close(messages)
}

func main() {
	emails := make(chan string, 2)
	calls := make(chan string, 2)
	messages := make(chan string, 3)

	go emailManager(emails)
	go phoneManager(calls)
	go instantMessages(messages)

	fmt.Println("🧑‍💼 Personal Assistant starting work day...")

	// Handle communications as they come in
	for {
		select {
		case email, ok := <-emails:
			if !ok {
				emails = nil // Channel closed, don't select it again
			} else {
				fmt.Printf("📧 Assistant: Handling %s\n", email)
			}
		case call, ok := <-calls:
			if !ok {
				calls = nil // Channel closed, don't select it again
			} else {
				fmt.Printf("📞 Assistant: Answering %s\n", call)
			}
		case msg, ok := <-messages:
			if !ok {
				messages = nil // Channel closed, don't select it again
			} else {
				fmt.Printf("💬 Assistant: Responding to %s\n", msg)
			}
		}

		// Exit when all channels are closed
		if emails == nil && calls == nil && messages == nil {
			break
		}

		time.Sleep(500 * time.Millisecond) // Small delay between tasks
	}

	fmt.Println("🧑‍💼 Assistant: All communications handled, end of workday!")
}
