package main

import "fmt"

// l -> r -> ... (adding 1 in each chain)
func main() {
	n := 10000
	leftmost := make(chan int)

	left := leftmost
	for range n {
		right := make(chan int) // need a new left link everytime
		go func(right, left chan int) {
			right <- 1 + <-left // the left chain will block until it receives data
		}(right, left)
		left = right // updating left as new right
	}

	// invoke the chain
	go func(c chan int) {
		c <- 1
	}(leftmost)

	fmt.Println(<-left) // result
}
