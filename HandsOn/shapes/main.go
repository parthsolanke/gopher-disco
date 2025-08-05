package main

import "fmt"

//Area and Perimeter of a Rectangle
//Problem:
//Define a struct Rectangle with fields Width and Height (both float64).
//Add two methods:
//* Area() float64
//* Perimeter() float64
//Test these methods in the main function.

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	r1 := Rectangle{Width: 12, Height: 13}
	fmt.Println(r1.Area())
	fmt.Println(r1.Perimeter())
}
