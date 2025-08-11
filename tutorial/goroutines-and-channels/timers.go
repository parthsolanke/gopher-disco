package main

import (
	"fmt"
	"time"
)

// Timers are used for when you want to do something once in teh future
// Timers will represent a single event in the future.
// You tell the timer how long you want to wait, and it provides a channel that will be notified at that time.
func main() {

	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
