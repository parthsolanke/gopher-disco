package main

import "fmt"

func sum(nums ...int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}
	return totalSum
}

func printType(values ...interface{}) {
	for _, val := range values {
		fmt.Printf("Value: %v, is of type: %T\n", val, val)
	}
}

func main() {
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sum([]int{1, 2, 3, 4, 5}...))
	printType("hello", 42, 3.14, true)
}
