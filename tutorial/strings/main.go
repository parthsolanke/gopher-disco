package main

import "fmt"

func analyzeText() {
	text := "Hello, 世界!"

	// Iterate over runes (Unicode code points)
	for index, runeValue := range text {
		fmt.Printf("Index %d: Rune %c (Unicode: %U)\n",
			index, runeValue, runeValue)
	}
}

func main() {
	// a rune is a 32bit integer
	analyzeText()
}
