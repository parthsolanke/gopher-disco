package main

import (
	"fmt"
	"time"
)

func slowBarista(coffee chan string) {
	time.Sleep(3 * time.Second) // Very slow barista
	coffee <- "☕ Your latte is ready!"
}

func main() {
	coffeeOrder := make(chan string)

	go slowBarista(coffeeOrder)

	fmt.Println("☕ Customer: I'd like a latte, please!")

	select {
	case drink := <-coffeeOrder:
		fmt.Printf("😊 Customer: Great! %s\n", drink)
	default:
		fmt.Println("😤 Customer: This is taking too long, I'm going to another café!")
	}
}
