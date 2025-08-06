package main

import "fmt"

// clarity: a pointer is nothing but a memory location
// &i gives you the address (pointer) to the var
// *T will be the type of var you store the address (pointer) eg: var ptr *int = &i
// *ptr is used to dereference the pointer and access the value of it
func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	fmt.Println("\nPointers part II: ")
	// pointer syntax explanation
	value := 5
	p := &value // p is a pointer to value
	fmt.Printf("Type of p: %T\n", p)
	fmt.Println(*p)    // prints 5 (the value at the address p points to)
	*p = 10            // changes value to 10
	fmt.Println(value) // prints 10

	// zero value pointers dont pint to anything
	var pt *int
	fmt.Println(pt == nil) // prints true

	// init pointers without zero values:
	// 1. using new keyword to create a pointer to a zero valued object
	r := new(int)
	fmt.Println(r)
	// 2. assigning it the address of an existing variable: p = &someInt

	// pointers in loops:
	// WRONG - all pointers will point to the last iteration value
	var ptrs1 []*int
	for i := 0; i < 3; i++ {
		ptrs1 = append(ptrs1, &i) // Bug: all point to the same variable
	}
	// above happen because at the beginning is initialized
	// same var is used at each iteration value change same memory is used

	// CORRECT - create a new variable in each iteration
	var ptrs2 []*int
	for i := 0; i < 3; i++ {
		val := i
		ptrs2 = append(ptrs2, &val)
	}
}
