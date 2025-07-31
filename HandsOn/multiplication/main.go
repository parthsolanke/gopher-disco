// Multiplication Table
// Print a 10x10 multiplication table using nested for loops.
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, val1 := range nums {
		for _, val2 := range nums {
			fmt.Print(val1 * val2)
		}
		fmt.Println()
	}
}
