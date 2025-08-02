package main

import "fmt"

// Write a function GroupByRemainder(nums []int, n int) map[int][]int
// that groups numbers by their remainder when divided by n.

func GroupByRemainder(nums []int, n int) map[int][]int {
	numsMap := make(map[int][]int)
	for _, val := range nums {
		rem := val % n
		numsMap[rem] = append(numsMap[rem], val)
	}

	return numsMap
}

func main() {
	fmt.Println(GroupByRemainder([]int{1, 2, 3, 4, 5, 6}, 2))
}
