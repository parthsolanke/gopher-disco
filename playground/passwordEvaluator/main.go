package main

import (
	"bufio"
	"fmt"
	"os"
	"passwordEvaluator/core"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter password: ")

	if scanner.Scan() {
		core.Evaluate(scanner.Text())
	} else {
		fmt.Println("Error while taking input")
	}
}
