package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "  Hello, World!  "

	// Common operations
	fmt.Println(strings.TrimSpace(text))         // "Hello, World!"
	fmt.Println(strings.ToUpper(text))           // "  HELLO, WORLD!  "
	fmt.Println(strings.ToLower(text))           // "  hello, world!  "
	fmt.Println(strings.Contains(text, "World")) // true

	// Splitting and joining
	parts := strings.Split("apple,banana,cherry", ",")
	fmt.Println(parts) // [apple banana cherry]

	joined := strings.Join([]string{"red", "green", "blue"}, "-")
	fmt.Println(joined) // "red-green-blue"
}
