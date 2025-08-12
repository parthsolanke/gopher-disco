package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
)

type Item struct {
	Price    float64
	Quantity int
}

type Stock struct {
	items map[string]*Item
	mu    sync.Mutex
}

func getStock() *Stock {
	return &Stock{
		items: map[string]*Item{
			"phone":   {Price: 699.99, Quantity: 5},
			"charger": {Price: 29.99, Quantity: 10},
			"laptop":  {Price: 1299.99, Quantity: 2},
			"mouse":   {Price: 19.99, Quantity: 5},
			"tablet":  {Price: 499.99, Quantity: 3},
		},
	}
}

func (s *Stock) String() string {
	s.mu.Lock()
	defer s.mu.Unlock()

	var b strings.Builder
	b.WriteString("Available Stock:\n")
	for name, item := range s.items {
		_, err := fmt.Fprintf(&b, "%s- price: %.2f, quantity: %d\n", name, item.Price, item.Quantity)
		if err != nil {
			fmt.Println("no stock available")
		}
	}
	return b.String()
}

func calculateOrderAmount(stock *Stock, items []string) float64 {
	stock.mu.Lock()
	defer stock.mu.Unlock()

	total := 0.0
	for _, name := range items {
		if stockItem, exists := stock.items[name]; exists {
			total += stockItem.Price
		}
	}
	return total
}

func processPayment(amount float64) error {
	if rand.Float64() < 0.3 {
		return fmt.Errorf("payment timeout")
	}
	return nil
}

func (s *Stock) deductStock(items []string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, name := range items {
		if stockItem, exists := s.items[name]; !exists || stockItem.Quantity <= 0 {
			return fmt.Errorf("%s out of stock", name)
		}
		s.items[name].Quantity--
	}

	return nil
}
