package main

import "fmt"

// composition is different than embedding in Go
// composition is mainly used to create complex types in Go
// using composition methods aren't promoted, only fields
// composition is a "has-a" relationship

type Car struct {
	Model string
}

func (c Car) describe() {
	fmt.Println("Car model:", c.Model)
}

// Truck composes Car (has-a Car), but does not embed it
type Truck struct {
	Car     Car // composed field, not embedded
	BedSize int
}

func main() {
	// Create a composed Truck
	t := Truck{
		Car:     Car{Model: "Toyota"},
		BedSize: 8,
	}

	// Access composed field explicitly
	fmt.Println(t.Car.Model) // Must go through Car

	// Methods are not promoted in composition
	// t.describe()           // ❌ Compile error: Truck has no describe() method

	// Must call method through composed field
	t.Car.describe()

	// Interface example
	type describer interface {
		describe()
	}

	var d describer
	// d = t       // ❌ Compile error: Truck does not implement describer
	d = t.Car // ✅ Car implements describer
	d.describe()
}
