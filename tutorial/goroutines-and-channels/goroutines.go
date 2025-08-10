package main

// A goroutine is a lightweight thread of execution.
import (
	"fmt"
	"time"
)

// concurrency is fake parallelism
// concurrency is where multiple tasks by rapidly switching between them (like task juggling)
// goroutines exhibit concurrency
func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// normal function call
	f("direct")

	// following goroutines a run concurrently
	go f("goroutine")

	// you can also start goroutine using anonymous function
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}

// Out:
// When we run this program, we see the output of the blocking call first, then the output of the two goroutines.
// The goroutines’ output may be interleaved because goroutines are being run concurrently by the Go runtime.
