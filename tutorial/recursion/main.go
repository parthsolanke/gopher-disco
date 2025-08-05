package main

import "fmt"

// eg : recursion
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(5))

	// fun signature
	var fib func(num int) int

	// assigning value to our func signature using sa anonymous function
	fib = func(num int) int {
		if num < 2 {
			return 1
		}
		return fib(num-1) + fib(num-2)
	}

	fmt.Println(fib(4))
}
