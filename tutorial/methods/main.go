package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

// method bound to the type rectangle
// r is a receiver similar to this in other languages,
// and Rectangle is the receiver type
// the receiver copies ove the struct hence inefficient for huge structs
func (r Rectangle) area() float64 {
	return r.Width * r.Height
}

// method with a pointer receiver (efficient for large structs)
// use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.
func (r Rectangle) perim() float64 {
	// dot can be used on the pointers to access the fields
	// it an automatically dereferences the pointer
	return 2 * (r.Width + r.Height)
}

func main() {
	r1 := Rectangle{Width: 12, Height: 16}
	fmt.Println(r1.area())
	// Go automatically handles conversion between values and pointers for method calls
	fmt.Println(r1.perim())

	// Go automatically handles conversion between pointers and values for method calls
	rp := &r1
	fmt.Println(rp.area())
	fmt.Println(rp.perim())
}
