package main

import "fmt"

func mergeMaps(map1, map2 map[string]int) {
	res := make(map[string]int)
	for key := range map1 {
		_, ok := map2[key]
		if ok {
			res[key] = map1[key] + map2[key]
		} else {
			res[key] = map1[key]
		}
	}
	fmt.Println(res)
}

func main() {
	// Problem:
	// Write a function that merges two map[string]int by adding the values for common keys.
	// If a key only exists in one map, copy it as is
	map1 := map[string]int{
		"parth": 1,
		"bob":   1,
		"jhon":  1,
	}
	map2 := map[string]int{
		"parth": 1,
		"bob":   1,
	}

	mergeMaps(map1, map2)
}
