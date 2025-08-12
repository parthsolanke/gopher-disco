package main

import "fmt"

// generic function for swap
func swap[T any](a, b T) (T, T) {
	return b, a
}

// here following is a generic function which accepts values of type T (int or float64)
func add[T int | float64](a, b T) T {
	return a + b
}

// generic function only accepting types which are comparable
func isEqual[T comparable](a, b T) bool {
	return a == b
}

// Not comparable (result in compilation error)
// Slices ([]int) ❌
// Maps (map[string]int) ❌
// Functions ❌
// Channels ❌
// Structs with non-comparable fields ❌

func Index[T comparable](s []T, val T) int {
	for i := range s {
		if s[i] == val {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(add(1, 2))     // int
	fmt.Println(add(3.5, 2.5)) // float64

	x, y := swap(1, 2)
	fmt.Println(x, y) // prints: 2 1

	// can also explicitly specify type (optimal)
	s1, s2 := swap[string]("hello", "world")
	fmt.Println(s1, s2) // prints: world hello

	fmt.Println(Index([]int{1, 2, 3}, 3))       // prints 1
	fmt.Println(Index([]string{"a", "b"}, "b")) // prints 1
}
