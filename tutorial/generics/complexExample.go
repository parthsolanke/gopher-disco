package main

import "fmt"

// custom constraint
// the tilde(~) in terms of generics is used to specify any type
// or custom type like "type Age T" whcih is strictly comprising of Tis allowed
// if ~ not included teh it wont accept custom types
// the "|" is for or
// thsi syntax is specific to generics
// otherwise ~ is for bitwise not in go
type Number interface {
	~int | ~float64
}

func Add[T Number](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(Add(1, 2))     // int
	fmt.Println(Add(3.5, 2.5)) // float64
}
