package main

import "fmt"

func doubleSlice(s []int) {
	for i, val := range s {
		s[i] = val * 2
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	doubleSlice(s)
	fmt.Println(s)
}
