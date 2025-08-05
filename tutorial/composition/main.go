package main

import "fmt"

// Go encourages composition through a feature called struct embedding, which is Go's alternative to inheritance.
type Car struct {
	Model string
}

type Truck struct {
	Car     // Embedded struct
	BedSize int
}

func main() {
	// When you embed Car into Truck, all of Car's fields are accessible directly from Truck
	t := Truck{
		Car:     Car{Model: "Ford"},
		BedSize: 10,
	}

	fmt.Println(t.Model)     // Accesses embedded field directly
	fmt.Println(t.Car.Model) // Accesses embedded field nested
	fmt.Println(t.BedSize)   // Truck's own field
}
