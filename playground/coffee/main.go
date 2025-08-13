package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func barista(orders <-chan int, serve chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for orderId := range orders {
		fmt.Println("processing order: ", orderId)
		time.Sleep(time.Duration(rand.Intn(5e3)) * time.Millisecond) // make coffee
		fmt.Printf("\norder %d complete\n", orderId)
		serve <- orderId
	}
}

func main() {
	var wg sync.WaitGroup
	ordersCh := make(chan int)
	serveCh := make(chan int)
	numCustomers := 5

	// ready baristas
	// starts 5 baristas
	for range 5 {
		wg.Add(1)
		go barista(ordersCh, serveCh, &wg)
	}

	// send orders
	// sends orders one by one so main doesn't block here
	go func() {
		for orderId := range numCustomers {
			ordersCh <- orderId
		}
		// close orders so barista won't get blocked
		close(ordersCh)
	}()

	// wait for each order to complete and close orders channel
	// routine waits to complete orders hence main won't block
	go func() {
		wg.Wait()
		close(serveCh)
	}()

	// collect all orders (blocks main here)
	for completedOrderId := range serveCh {
		fmt.Println("received order: ", completedOrderId)
	}
}
