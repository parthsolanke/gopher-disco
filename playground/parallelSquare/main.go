package main

import (
	"fmt"
	"math"
)

func square(num int, ch chan int) {
	ch <- int(math.Pow(float64(num), 2))
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	ch := make(chan int)

	for _, num := range nums {
		go square(num, ch)
	}

	for range nums {
		fmt.Println(<-ch)
	}
}
