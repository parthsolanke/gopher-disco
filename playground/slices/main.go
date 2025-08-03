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
