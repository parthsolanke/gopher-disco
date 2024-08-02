package main

import "fmt"

func main() {
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Start of week")
	case "Friday":
		fmt.Println("End of work week")
	default:
		fmt.Println("Stilll work week")
	}
}
