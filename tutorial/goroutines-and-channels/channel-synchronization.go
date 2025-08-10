package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Send a value to notify that we’re done.
	done <- true
}

func main() {

	done := make(chan bool, 1)
	// Start a worker goroutine, giving it the channel to notify on.
	go worker(done)

	// Block until we receive a notification from the worker on the channel.
	<-done
	// If you removed the <- done line from this program, the program would exit before the worker even started.
}
