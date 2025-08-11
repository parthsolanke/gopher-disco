package main

import "fmt"

// Closing a channel indicates that no more values will be sent on it.
// This can be useful to communicate completion to the channel’s receivers.
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			if job, more := <-jobs; more {
				fmt.Printf("Recieved job: %d\n", job)
			} else {
				fmt.Println("Done receiving jobs")
				done <- true
				return
			}
		}
	}()

	for i := range 3 {
		jobs <- i
		fmt.Printf("Sent job: %d\n", i)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done

	_, ok := <-jobs
	fmt.Println("more jobs: ", ok)
}
