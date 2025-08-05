package main

import (
	"fmt"
	"math"
)

// polymorphism using interfaces
type shape interface {
	area() float64
	perim() float64
}

type rectangle struct {
	width  float64
	height float64
}

func (r *rectangle) area() float64 {
	return r.width * r.height
}

func (r *rectangle) perim() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	Radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.Radius
}

func getshapeInfo(s shape) {
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perim())
}

func detectCircle(g shape) {
	// using type assertion to get the type of the shape
	if c, ok := g.(Circle); ok {
		fmt.Println("circle with radius", c.Radius)
	} else {
		fmt.Println("Not a circle")
	}
}

// compiler check if a type implements a certain interface or not
// this line will cause a compiler error, catching the mistake early rather than at runtime.
var _ shape = (*Circle)(nil)

func main() {
	c1 := Circle{Radius: 12}
	getshapeInfo(c1)

	r1 := rectangle{height: 12, width: 13}
	// in case of interfaces go doesn't automatically take address hence have to be explicit
	getshapeInfo(&r1) // pass pointer

	// run time type checking in golang
	detectCircle(c1)
	detectCircle(&r1)
}
