package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// non-blocking receives
	// as no value available on the chan
	//  the default case will execute
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// non-blocking sends
	// no receiver on the can hence default
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// both of the above cases
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
