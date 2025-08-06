package main

import "fmt"

type Car struct {
	Model string
}

func (c Car) describe() {
	fmt.Println("Car model: ", c.Model)
}

// struct embedding, which is Go's alternative to inheritance.
// using embedding in go methods are also promoted

type Truck struct {
	Car     // Embedded struct, fields in Car now are the part of Truck
	BedSize int
}

func main() {
	// When you embed Car into Truck, all of Car's fields are accessible directly from Truck
	t := Truck{
		Car:     Car{Model: "Ford"},
		BedSize: 10,
	}

	fmt.Println(t.Model)     // Accesses embedded field directly
	fmt.Println(t.Car.Model) // Alternatively, we can spell out the full path using the embedded type name
	t.describe()             // Since container embeds base, the methods of base have also become methods of a container.

	// embedding structs with methods are also used to bestow them to implement an interface
	type describer interface {
		describe()
	}

	// here we see that the truck now implements the describer interface because it embeds car
	d := t
	d.describe()

}
