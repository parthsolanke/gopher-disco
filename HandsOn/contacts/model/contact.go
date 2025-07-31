package model

import "fmt"

type Contact struct {
	Name     string
	Email    string
	Age      int
	Height   float64
	IsActive bool
}

func Summary(c Contact) string {
	return fmt.Sprintf("%+v", c)
}

func PrintDetails(c Contact) {
	fmt.Println("\nContact: ")
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Email: %s\n", c.Email)
	fmt.Printf("Age: %d\n", c.Age)
	fmt.Printf("Height: %.2f cm\n", c.Height)
	fmt.Printf("Active: %t\n", c.IsActive)
	fmt.Println("Summary:", Summary(c))
}
