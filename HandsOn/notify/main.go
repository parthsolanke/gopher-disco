package main

import "fmt"

//Create an interface Notifier with a method Notify().
//Implement it for two structs:
//* EmailUser with Email field
//* SMSUser with Phone field
//Each should print a different message when Notify() is called.
//Demonstrate polymorphism by calling Notify() on both using a slice of Notifier.

type Notifier interface {
	Notify()
}

type EmailUser struct {
	Email string
}

type SMSUser struct {
	Phone []int
}

func (e EmailUser) Notify() {
	fmt.Println("Notified from email")
}

func (e SMSUser) Notify() {
	fmt.Println("Notified from sms")
}

func getNotified(n Notifier) {
	n.Notify()
}

func main() {
	e1 := EmailUser{Email: "parth@gmail.com"}
	getNotified(e1)

	s1 := SMSUser{Phone: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}}
	getNotified(s1)

}
