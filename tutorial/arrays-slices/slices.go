package main

import (
	"fmt"
)

func main() {

	//slices
	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	// adding elements in a slice
	myslice2 = append(myslice2, "are they")
	fmt.Println(myslice2)

	// slices form arrays
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	s := [3]int{1, 2, 3}
	fmt.Println(len(s), cap(s))
	s1 := s[0:1]
	fmt.Println(len(s1), cap(s1))
	s1 = append(s1, 5)
	fmt.Println(s1)
	fmt.Println(s)

	// creating empty slices with specific size
	slice := make([]int, 3, 5)
	fmt.Println(slice)
	fmt.Println("Length:", len(slice))   // 3 (current items)
	fmt.Println("Capacity:", cap(slice)) // 5 (max before reallocation)
}

//The Slice Growth Formula
//When a slice runs out of capacity, Go doesn't just add one more spot - it follows a clever formula:
//
//If the slice has less than 1,024 elements, it doubles the capacity
//If it has more than 1,024 elements, it grows by 25%
//This balance prevents both frequent memory allocations (expensive) and excessive memory waste (also bad). It's like Go has a built-in economist deciding the perfect growth rate!
