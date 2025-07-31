package main

import (
	"expressions/core"
	"fmt"
)

func main() {
	num1, num2, err := core.ProcessIn()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Addition: %d\n", num1+num2)
	fmt.Printf("Substraction: %d\n", num1-num2)
	fmt.Printf("Multiplication: %d\n", num1*num2)
	fmt.Printf("Division: %.2f\n", float64(num1)/float64(num2))
	fmt.Printf("Mod: %d\n", num1%num2)

}
