package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

const STOCK_PRICE = 23.58

func getPriceService(ctx context.Context, name string) <-chan float64 {
	priceCh := make(chan float64)

	go func() {
		defer close(priceCh)
		delay := time.Duration(rand.Intn(1000)) * time.Millisecond
		select {
		case <-time.After(delay):
			priceCh <- STOCK_PRICE
		case <-ctx.Done():
			// Cancelled, don't send anything
			return
		}
	}()

	return priceCh
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure all goroutines are cleaned up

	api1 := getPriceService(ctx, "api1")
	api2 := getPriceService(ctx, "api2")
	api3 := getPriceService(ctx, "api3")

	select {
	case p := <-api1:
		fmt.Println("Data received from api1:", p)
		cancel()
	case p := <-api2:
		fmt.Println("Data received from api2:", p)
		cancel()
	case p := <-api3:
		fmt.Println("Data received from api3:", p)
		cancel()
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout")
	}
}
