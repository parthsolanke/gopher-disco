package main

import (
	"fmt"
	"strings"
)

func main() {
	//	Capitalize Each Word
	//	Given a sentence, capitalize the first letter of each word.
	//	Input: "hello world"
	// Output: "Hello World"
	text := "It's a wonderful day"

	words := strings.Split(text, " ")
	fmt.Println(words)

	res := []string{}
	for _, word := range words {
		temp := strings.ToUpper(string(word[0])) + word[1:]
		res = append(res, temp)
	}
	fmt.Println(strings.Join(res, " "))
}
