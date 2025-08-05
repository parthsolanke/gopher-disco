package main

import "fmt"

// Function using type assertion
func handleWithAssertion(x interface{}) {
	fmt.Println("Using Type Assertion:")

	// Try to assert x as a string
	if str, ok := x.(string); ok {
		fmt.Println("It's a string:", str)
	} else {
		fmt.Println("Type assertion to string failed!")
	}

	// Try to assert x as an int
	if i, ok := x.(int); ok {
		fmt.Println("It's an int:", i)
	} else {
		fmt.Println("Type assertion to int failed!")
	}

	fmt.Println()
}

// Function using type switch
func handleWithTypeSwitch(x interface{}) {
	fmt.Println("Using Type Switch:")

	switch v := x.(type) {
	case string:
		fmt.Println("It's a string:", v)
	case int:
		fmt.Println("It's an int:", v)
	case bool:
		fmt.Println("It's a bool:", v)
	default:
		fmt.Println("Unknown type")
	}

	fmt.Println()
}

func main() {
	var val1 interface{} = "Hello, Go!"
	var val2 interface{} = 42
	var val3 interface{} = true

	handleWithAssertion(val1)
	handleWithAssertion(val2)
	handleWithTypeSwitch(val1)
	handleWithTypeSwitch(val2)
	handleWithTypeSwitch(val3)
}
