package main

import "fmt"

// using for to iterate over channels
func main() {
	queue := make(chan int, 2)
	queue <- 1
	queue <- 2
	close(queue) // important to close a channel when using for over channels

	for e := range queue {
		fmt.Println(e)
	}
}
