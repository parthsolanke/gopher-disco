package main

import "fmt"

func totalCost(prices ...float64) (float64, int) {
	sum := 0.0
	for _, num := range prices {
		sum += num
	}
	return sum, len(prices)
}

func main() {
	total, n := totalCost([]float64{1, 2, 3, 45, 5}...)
	fmt.Printf("Total price: %.2f and total items: %d\n", total, n)
}
