package main

import (
	"fmt"
	"math/rand"
	"time"
)

// for is like “stay in the loop” command – literally instructs the goroutine to remain active.

func main() {
	c := make(chan string)
	go simpleBoring("boring! ", c)

	for range 5 {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You are boring; I'm leaving")
}

func simpleBoring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		// 10^3 * 10^6 (nanoseconds) = 1 second
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
