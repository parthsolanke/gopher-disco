package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
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

	// string formating
	name := "Alice"
	age := 30
	score := 95.7

	// Basic formatting
	fmt.Printf("Name: %s\n", name)     // Name: Alice
	fmt.Printf("Age: %d\n", age)       // Age: 30
	fmt.Printf("Score: %.1f\n", score) // Score: 95.7

	// Multiple values
	fmt.Printf("%s is %d years old and scored %.1f%%\n", name, age, score)
	// Output: Alice is 30 years old and scored 95.7%

	text = "Hello世界🚀"

	// Counting bytes (what len() does)
	fmt.Printf("Bytes: %d\n", len(text)) // 13 bytes

	// Counting runes (actual characters)
	runeCount := 0
	for range text {
		runeCount++
	}
	fmt.Printf("Characters: %d\n", runeCount) // 8 characters

	// Alternative way to count runes
	fmt.Printf("Runes: %d\n", len([]rune(text))) // 8 characters

	// printing
	for i, r := range text {
		fmt.Printf("%d: %c, type: %T\n", i, r, r) // Iterates runes, not bytes
	}

	// to upper case
	fmt.Println(strings.ToUpper(text))

	// vs
	runes := []rune(text)
	for i, r := range runes {
		runes[i] = unicode.ToUpper(r)
	}

	fmt.Println(string(runes))

	// sorting and searching in strings
	// Sort strings alphabetically
	names := []string{"Charlie", "Alice", "Bob"}
	sort.Strings(names)
	fmt.Println(names) // [Alice Bob Charlie]

	// custom sorting
	// Sort people by name length
	people := []struct {
		Name string
		Age  int
	}{
		{"Alice", 30},
		{"Bob", 25},
		{"Alexander", 35},
	}

	// Sort by name length (shortest first)
	sort.Slice(people, func(i, j int) bool {
		return len(people[i].Name) < len(people[j].Name)
	})

	// searching in string:
	text = "Hello, World!"
	if strings.Contains(text, "World") {
		fmt.Println("Found 'World' in text")
	}

	index := strings.Index(text, "World")
	if index != -1 {
		fmt.Printf("'World' starts at position %d\n", index)
	}
}
