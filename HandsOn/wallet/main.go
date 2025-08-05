package main

import "fmt"

//Create a struct Account with fields:
//* HolderName (string)
//* Balance (float64)
//Add methods:
//* Deposit(amount float64)
//* Withdraw(amount float64)
//* Display()
//Ensure Withdraw doesn't allow the balance to go negative.

type DebitCard struct {
	HolderName      string
	DebitCardNumber int
}

type Account struct {
	HolderName string
	Balance    float64
	DebitCard
}

func (a *Account) Deposit(amount float64) {
	fmt.Println()
	if amount > 0 {
		fmt.Printf("Old balance: %v\n", a.Balance)
		a.Balance += amount
		fmt.Printf("New balance: %v\n", a.Balance)
		return
	}
	fmt.Println("InvalDebitCardNumber amount")
}

func (a *Account) Display() {
	fmt.Println(a.HolderName)
	fmt.Println(a.Balance)
}

func (a *Account) Withdraw(amount float64) {
	fmt.Println()
	if num := a.Balance - amount; num > 0 && a.DebitCard.DebitCardNumber != 0 {
		fmt.Printf("Old balance: %v\n", a.Balance)
		a.Balance -= amount
		fmt.Printf("New balance: %v\n", a.Balance)
		return
	}
	fmt.Println("Balance insufficient")
}

func main() {
	a1 := Account{HolderName: "Parth",
		Balance: 1000,
		DebitCard: DebitCard{
			DebitCardNumber: 123,
			HolderName:      "Parth",
		}}
	a1.Display()
	a1.Deposit(200)
	a1.Withdraw(1000)
}
