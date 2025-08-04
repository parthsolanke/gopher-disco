package main

import (
	"fmt"
)

func invertMap(in map[interface{}]interface{}) {
	if len(in) == 0 {
		fmt.Println("Empty map")
	}

	res := make(map[interface{}]interface{})
	for key, val := range in {
		res[val] = key
	}

	fmt.Println(res)
}

func main() {
	//	Invert a Map
	//  Problem:
	//	Write a function that inverts a map[string]string so that the values become keys and the keys become values.

	myMap := map[string]int{
		"parth": 1,
		"bob":   2,
		"jhon":  3,
	}

	myMapInterface := make(map[interface{}]interface{})
	for k, v := range myMap {
		myMapInterface[k] = v
	}

	invertMap(myMapInterface)

}
