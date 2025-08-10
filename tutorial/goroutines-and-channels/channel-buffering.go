package main

// Buffered channels accept a limited number of values without a corresponding receiver for those values.
import "fmt"

func main() {
	// Here we make a channel of strings buffering up to 2 values.
	messages := make(chan string, 2)

	// Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.
	// context: the main thread is running sequentially
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

// Producer-Consumer: When producer and consumer work at different speeds
// Batch Processing: Collecting items before processing them together
// Rate Limiting: Controlling how many operations happen at once
