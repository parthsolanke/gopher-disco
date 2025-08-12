package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	stock := getStock()
	fmt.Println(stock)
	var numOrders int
	fmt.Print("Enter number of orders: ")

	if _, err := fmt.Scan(&numOrders); err != nil {
		fmt.Println("invalid input")
	}

	// take order inputs from user
	orders := make([]Order, 0, numOrders)
	reader := bufio.NewReader(os.Stdin)
	for i := range numOrders {
		orders = append(orders, takeOrderInput(i+1, reader, stock))
	}

	// worker pool
	numWorkers := 5
	orderChan := make(chan Order)
	errorsChan := make(chan error, numOrders)
	var wg sync.WaitGroup

	// start workers
	for i := range numWorkers {
		wg.Add(1)
		go worker(stock, i, orderChan, &wg, errorsChan)
	}

	// invoking workers by sending jobs / orders
	go func() {
		for _, order := range orders {
			orderChan <- order
		}
		close(orderChan)
	}()

	// wait for the workers to finish
	wg.Wait()
	close(errorsChan)

	// look at all teh errors
	fmt.Println("\n---errors---")
	for err := range errorsChan {
		fmt.Println("Error: ", err)
	}

}
