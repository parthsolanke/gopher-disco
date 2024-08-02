package main

import "fmt"

type Person struct {
	name    string
	age     int
	school  string
	married bool
}

func main() {
	// decalration
	var p1 Person

	// initialization
	p1.name = "Parth"
	p1.age = 22
	p1.school = "Stranford"
	p1.married = false

	fmt.Println(p1)

	//short variable declaration operator
	p2 := Person{
		name:    "Ambarish",
		age:     22,
		school:  "Stanford",
		married: false,
	}

	fmt.Println(p2)
}
