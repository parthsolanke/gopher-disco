package main

import (
	"fmt"
)

func main() {
	score := 67
	scoreRating(score)
}

func scoreRating(score int) {
	switch {
	case score >= 90:
		fmt.Println("Excellent")
	case score >= 75:
		fmt.Println("Good")
	case score >= 60:
		fmt.Println("Pass")
	default:
		fmt.Println("Fail")
	}
}
