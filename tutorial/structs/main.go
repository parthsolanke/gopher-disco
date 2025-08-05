package main

import "fmt"

// A and B are different type! as order of the fields is concerned
//type A struct {
//	X int
//	Y int
//}
//
//type B struct {
//	Y int
//	X int
//}

// structs are mutabel
type Person struct {
	name string
	age  int
}

// use pointers for efficiency (avoid copying over the whole struct)
// use pointers for mutability
func newPerson(name string) *Person {
	p := Person{name: name}
	p.age = 42
	return &p
}

func main() {
	// create new structs
	fmt.Println(Person{"parth", 70})

	// omitted fields will be zero-valued
	fmt.Println(Person{name: "Fred"})

	// An & prefix yields a pointer to the struct.
	fmt.Println(&Person{name: "Ann", age: 40}) // formats to human relabel way
	fmt.Printf("%p\n", &Person{name: "Ann", age: 40})

	// Access struct fields with a dot.
	s := Person{name: "Sean", age: 50}
	fmt.Println(s.name)

	//You can also use dots with struct pointers - the pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.age)

	// anonymous structs:
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
