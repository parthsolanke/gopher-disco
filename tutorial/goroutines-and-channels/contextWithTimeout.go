package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a context with a 2-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // always cancel to free resources

	// Simulate work in a goroutine
	go func() {
		// pretend we're doing something that takes 5 seconds
		fmt.Println("Work started...")
		time.Sleep(5 * time.Second)
		fmt.Println("Work finished!") // this will never be printed
	}()

	// Wait until either work finishes or timeout occurs
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Work completed before timeout")
	case <-ctx.Done():
		// ctx.Err() will be context.DeadlineExceeded
		// can do return here if in a function basically exit
		fmt.Println("Timeout reached:", ctx.Err())
	}
}
