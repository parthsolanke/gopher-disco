package main

import (
	"errors"
	"fmt"
	"time"
)

// the entire error sys in go concocts on one singed interface
// any struct can implement this interface
// error ca be of any type
// type error interface {
//	Error() string
//}

// --- Custom Error Type ---
type ValidationError struct {
	Field     string
	Value     interface{}
	Reason    string
	Timestamp time.Time
}

// implementing teh error interface
func (ve ValidationError) Error() string {
	return fmt.Sprintf("validation failed for '%s': %s (value: %v)",
		ve.Field, ve.Reason, ve.Value)
}

func (ve ValidationError) IsTemporary() bool { return false }

// --- Example Functions ---

// Custom error example
func validateAge(age int) error {
	if age < 0 {
		return ValidationError{"age", age, "negative value", time.Now()}
	}
	if age > 150 {
		return ValidationError{"age", age, "unrealistic value", time.Now()}
	}
	return nil
}

// Simple error example
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero") // Quick error
	}
	return a / b, nil
}

// Formatted error example
func withdraw(balance, amount float64) (float64, error) {
	if amount > balance {
		return balance, fmt.Errorf("insufficient funds: tried $%.2f but only have $%.2f",
			amount, balance) // Detailed error
	}
	return balance - amount, nil
}

// --- Main ---
func main() {
	// 1️⃣ Quick error
	if _, err := divide(10, 0); err != nil {
		fmt.Println("❌", err)
	}

	// 2️⃣ Formatted error
	if _, err := withdraw(100, 150); err != nil {
		fmt.Println("❌", err)
	}

	// 3️⃣ Custom error
	if err := validateAge(-5); err != nil {
		fmt.Println("❌", err)
		// we need to assert as the return type is error interface
		// if we don't assert then go will panic
		if ve, ok := err.(ValidationError); ok {
			fmt.Println("🕐 Occurred at:", ve.Timestamp)
		}
	}
}
