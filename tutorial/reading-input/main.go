package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Accessing command-line arguments (excluding the program name)
	args := os.Args[1:]

	if len(args) > 0 {
		fmt.Println("Command-line arguments:")
		for i, arg := range args {
			fmt.Printf("Arg %d: %s\n", i+1, arg)
		}
	} else {
		fmt.Println("No command-line arguments provided.")
	}

	// Taking input from the user via stdin
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a number: ")

	if scanner.Scan() {
		input := scanner.Text()
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid number:", err)
			return
		}
		fmt.Println("You entered:", number)
	} else {
		fmt.Println("Failed to read input.")
	}
}
