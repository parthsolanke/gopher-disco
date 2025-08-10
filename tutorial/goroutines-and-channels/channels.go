package main

import "fmt"

func main() {
	// Create a channel that passes string values
	// also this is a type fo unbuffered channel
	// An unbuffered channel has no capacity to store values.
	// A send operation (ch <- val) will block until another goroutine is ready to receive (<- ch).
	// A receive operation will block until another goroutine sends.
	messages := make(chan string)

	// Start a goroutine (a lightweight thread)
	// It runs this function: send "ping" into the channel
	//  (chan <-) is used to send msg form one go routine to another
	go func() {
		messages <- "ping" // <- means "send this into the channel"
	}()

	// Wait for the message to arrive and store it in msg
	// (<- chan) is used to get msg from a channel
	msg := <-messages

	// Print the message
	fmt.Println(msg)
}

// Synchronization: When you need perfect timing between goroutines
// Request-Response: Web servers handling HTTP requests
// Pipeline Processing: When each step must complete before the next begins
