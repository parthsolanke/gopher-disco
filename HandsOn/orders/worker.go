package main

import (
	"fmt"
	"strings"
	"sync"
)

func worker(stock *Stock, id int, orders <-chan Order, wg *sync.WaitGroup, errorCh chan<- error) {
	defer wg.Done()
	fmt.Println()
	for order := range orders {
		fmt.Printf("worker-%d processing order: %d, ", id, order.ID)

		// check stock
		if err := stock.deductStock(order.Items); err != nil {
			fmt.Println("out of Stock")
			errorCh <- fmt.Errorf("item %s out of stock for order %d", strings.Join(order.Items, ","), order.ID)
			continue
		}

		// process payment
		if err := processPayment(order.Amount); err != nil {
			fmt.Println("payment Failed")
			errorCh <- fmt.Errorf("payment failed for order %d: %v", order.ID, err)
			continue
		}

		fmt.Println("payment Success")
	}
}
