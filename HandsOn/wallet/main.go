package main

import (
	"errors"
	"fmt"
)

// DebitCard holds card details. Separate from Account logic.
type DebitCard struct {
	DebitCardNumber int
}

// Account represents a bank account with deposit/withdraw capabilities.
type Account struct {
	holderName string
	balance    float64
	card       *DebitCard // Pointer to allow optional association
}

// NewAccount is a constructor function for Account
func NewAccount(holderName string, initialBalance float64, card *DebitCard) *Account {
	return &Account{
		holderName: holderName,
		balance:    initialBalance,
		card:       card,
	}
}

// Deposit adds funds to the account.
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	fmt.Printf("Depositing %.2f to %s's account\n", amount, a.holderName)
	a.balance += amount
	return nil
}

// Withdraw removes funds if sufficient balance exists.
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdraw amount must be positive")
	}
	if a.card == nil || a.card.DebitCardNumber == 0 {
		return errors.New("valid debit card required for withdrawal")
	}
	if amount > a.balance {
		return errors.New("insufficient balance")
	}
	a.balance -= amount
	fmt.Printf("Withdrew %.2f from %s's account\n", amount, a.holderName)
	return nil
}

// Display shows account information.
func (a *Account) Display() {
	fmt.Println("=== Account Details ===")
	fmt.Printf("Holder: %s\n", a.holderName)
	fmt.Printf("Balance: %.2f\n", a.balance)
	if a.card != nil {
		fmt.Printf("Debit Card #: %d\n", a.card.DebitCardNumber)
	}
	fmt.Println("========================")
}

// Main execution point
func main() {
	card := &DebitCard{
		DebitCardNumber: 12345678,
	}
	account := NewAccount("Parth", 1000.0, card)

	account.Display()

	if err := account.Deposit(200); err != nil {
		fmt.Println("Deposit Error:", err)
	}
	if err := account.Withdraw(1000); err != nil {
		fmt.Println("Withdraw Error:", err)
	}

	account.Display()
}
