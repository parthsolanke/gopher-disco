package main

import "fmt"

func main() {
	fmt.Println(average(1, 2, 3, 4, 5, 6))
	fmt.Println(average([]float64{1, 3, 4, 5}...))
}

func average(nums ...float64) float64 {
	sum := 0.0
	for _, val := range nums {
		sum += val
	}
	return sum / float64(len(nums))
}
