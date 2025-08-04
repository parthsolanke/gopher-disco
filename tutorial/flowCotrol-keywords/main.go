package main

import "fmt"

// labeled break and continue
func findInMatrix(matrix [][]int, target int) (int, int) {
OuterLoop:
	for i, row := range matrix {
		for j, value := range row {
			if value == target {
				fmt.Printf("Found %d at position (%d, %d)\n", target, i, j)
				break OuterLoop // Break out of both loops
			}
		}
	}
	return -1, -1 // Not found
}

func processWithSkip() {
RowLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == j {
				continue RowLoop // Skip to next iteration of outer loop
			}
			fmt.Printf("(%d, %d) ", i, j)
		}
		fmt.Println()
	}
}

func main() {
	processWithSkip()
	findInMatrix([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{5, 6, 7},
	},
		3)
}
