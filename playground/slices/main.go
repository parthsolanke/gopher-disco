package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	printHeader("s1", s1)

	s2 := s1[:2]
	printHeader("s2", s2)

	s1 = s1[1:4]
	printHeader("reassigned s1", s1)
	// printing indexes from reassigned s1 which arent supposed to be accessible
	fmt.Println(s1[:7])
}

func printHeader(name string, s []int) {
	ptr := &s[0]
	fmt.Printf("%s -> slice: %v, addr: %p, len: %d, cap: %d\n", name, s, ptr, len(s), cap(s))
	fmt.Println()
}

// Out:
//~/dev/gopher-disco/playground/slices (main ✔) go run main.go
//s1 -> slice: [1 2 3 4 5 6 7 8], addr: 0xc00001a180, len: 8, cap: 8
//
//s2 -> slice: [1 2], addr: 0xc00001a180, len: 2, cap: 8
//
//reassigned s1 -> slice: [2 3 4], addr: 0xc00001a188, len: 3, cap: 7
//
//[2 3 4 5 6 7 8]
//~/dev/gopher-disco/playground/slices (main ✗)

// practice to avoid changing the og array:
//original := []int{1, 2, 3, 4, 5}
//copySlice := make([]int, 3)
//copy(copySlice, original[1:4]) // Copy [2, 3, 4] into new slice
//
//copySlice[0] = 99
//fmt.Println(original)   // [1, 2, 3, 4, 5]
//fmt.Println(copySlice)  // [99, 3, 4]
