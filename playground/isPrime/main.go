package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := range int(math.Sqrt(float64(num)) - 1) {
		i += 2
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter number to check: ")
		if !scanner.Scan() {
			break
		}

		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid number")
			continue
		}

		if isPrime(num) {
			fmt.Println("Prime number")
		} else {
			fmt.Println("Not prime")
		}
	}
}
