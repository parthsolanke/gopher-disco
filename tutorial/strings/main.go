package main

import "fmt"

func analyzeText() {
	text := "Hello, 世界!"

	// Iterate over runes (Unicode code points)
	for index, runeValue := range text {
		fmt.Printf("Index %d: Rune %c (Unicode: %U)\n",
			index, runeValue, runeValue)
	}

	//  len of teh string returns total bytes not chars
	fmt.Printf("Length of teh string: %d\n", len(text))
	// strings are sequence of bytes
	// normal chars are 1 byte each eg: "H", "E"
	// special chars are unicode-8 and occupy more space like 3 bytes each or more

	// getting all the chars
	fmt.Printf("len by chars: %d\n", len([]rune(text)))

	// getting teh first char its not string its a byte ASCII
	fmt.Printf("Value: %v, Type: %T\n", text[0], text[0])

	// converting to byte
	fmt.Println(string(text[0]))

}

func main() {
	analyzeText()
}
