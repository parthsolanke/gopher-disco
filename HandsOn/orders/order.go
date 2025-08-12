package main

import (
	"bufio"
	"fmt"
	"strings"
)

type Order struct {
	ID     int
	Items  []string
	Amount float64
}

func takeOrderInput(id int, reader *bufio.Reader, stock *Stock) Order {
	var order Order
	order.ID = id

	fmt.Printf("\n--- Order #%d ---\n", id)
	fmt.Print("Items (comma separated): ")
	itemsStr := readLine(reader)
	order.Items = strings.Split(itemsStr, ",")
	for i := range order.Items {
		order.Items[i] = strings.TrimSpace(order.Items[i])
	}

	order.Amount = calculateOrderAmount(stock, order.Items)
	return order
}
