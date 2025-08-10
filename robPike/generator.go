package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func main() {
	quit := make(chan bool)
	c := fanIn(boring("joe", quit), boring("ann", quit))
	timeout := time.After(5 * time.Second)

	for {
		select {
		case msg := <-c:
			fmt.Println(msg.str)
			msg.wait <- true
		case <-timeout:
			fmt.Println("Timeout! Quitting...")
			close(quit)
			return
		}
	}
}

// generator function
// same as foundations but with much better pattern
// generator acts as its own independent service
// returns receive only (a directional) channel
func boring(msg string, quit chan bool) <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool)

	go func() {
		for i := 0; ; i++ {
			select {
			case c <- Message{fmt.Sprintf("%s %d", msg, i), waitForIt}:
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
				<-waitForIt // blocks until main sends data (for sequential communication)
			case <-quit:
				fmt.Printf("%s: quitting\n", msg)
				return
			}
		}
	}()

	return c
}

// fan-in function
// used to let whoever is ready to first talk
// it will take values from 2 routines adn forward it to another which we will use in main
func fanIn(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)

	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()

	return c
}
