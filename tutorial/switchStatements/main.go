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

	whatAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("I am an int")
		case string:
			fmt.Println("I am a string")
		case bool:
			fmt.Println("I am a bool")
		default:
			fmt.Println("I am something else")
		}
	}
	whatAmI(42)
	whatAmI("Hello")
	whatAmI(true)
	whatAmI(3.14)
	whatAmI([]int{1, 2, 3})

	// using switch case to avoid long if-else chains
	number := 15
	switch {
	case number < 10:
		fmt.Println("Number is less than 10")
	case number < 20:
		fmt.Println("Number is less than 20")
	default:
		fmt.Println("Number is 20 or more")
	}

	// using fallthrough
	switch {
	case number < 10:
		fmt.Println("Number is less than 10")
		fallthrough
	case number < 20:
		fmt.Println("Number is less than 20")
		fallthrough
	case number < 30:
		fmt.Println("Number is less than 30")
	default:
		fmt.Println("Number is 30 or more")
	}

}
