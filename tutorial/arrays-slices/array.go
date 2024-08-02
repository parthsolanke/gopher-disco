package main

import "fmt"

func main() {
	//1.
	// Declaration
	var arr1 [5]string

	// Initialization
	arr1 = [5]string{"Volvo", "BMW", "Ford", "Mazda", "VW"}

	fmt.Println(arr1)

	//2.
	// Declaration + Initialization
	var arr2 [3]int = [3]int{1, 2, 3}

	fmt.Println(arr2)

	//3.
	// Declaration + Initialization using short variable declaration operator
	arr3 := [4]float64{1.24, 2.35, 9}

	fmt.Println(arr3)

	//4.
	// Arrays with inferred length
	arr4 := [...]bool{true, false, true}

	fmt.Println(arr4)
}
