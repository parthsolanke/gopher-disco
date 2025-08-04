package main

import "fmt"

func demonstrateOrder() {
	defer fmt.Println("First deferred")  // 1.
	defer fmt.Println("Second deferred") // 3.
	defer fmt.Println("Third deferred")  // 2.

	fmt.Println("Function body") // 1.
}

func main() {
	demonstrateOrder()
}
