package main

import "fmt"

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
}
