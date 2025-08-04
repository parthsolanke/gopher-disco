package main

import "fmt"

func main() {
	// creating maps
	myMap := make(map[string]int)
	fmt.Println(myMap)

	// creating maps wth initial values
	myMap2 := map[string][]int{
		"parth": {1, 2, 3, 4},
		"bob":   {1, 2, 3, 4},
	}
	fmt.Println(myMap2)
	val, ok := myMap2["part"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("value docent exist")
	}
	// getting len
	fmt.Println(len(myMap2))

	myMap2["jhon"] = []int{1, 3, 4}

	// deleting entry from map
	delete(myMap2, "jhon")
	fmt.Println(myMap2)

	// clearing teh whole map
	clear(myMap2)
	fmt.Println(myMap2)

	//Valid Map Keys
	//Keys must be comparable (you can use == with them):
	//
	//// ✅ Good keys
	//map[string]int{}        // strings
	//map[int]string{}        // numbers
	//map[bool]string{}       // booleans
	//
	//// ❌ Bad keys - these won't compile
	//map[[]int]string{}      // slices
	//map[map[string]int]string{} // maps

}
