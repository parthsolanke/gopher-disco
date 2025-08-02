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
}
