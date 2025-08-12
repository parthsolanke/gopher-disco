package main

import (
	"errors"
	"fmt"
)

func connectToDatabase() error {
	return errors.New("connection refused")
}

func getUserData(userID int) error {
	if err := connectToDatabase(); err != nil {
		// Wrap the error with more context - like leaving breadcrumbs
		return fmt.Errorf("failed to get user %d: %w", userID, err)
	}
	return nil
}

func handleRequest(userID int) error {
	if err := getUserData(userID); err != nil {
		// Add even more context
		return fmt.Errorf("request handling failed: %w", err)
	}
	return nil
}

func main() {
	if err := handleRequest(123); err != nil {
		fmt.Printf("❌ %v\n", err)
		// Output: request handling failed: failed to get user 123: connection refused

		// You can unwrap to find the root cause
		if errors.Is(err, errors.New("connection refused")) {
			fmt.Println("🔍 Root cause: Database connection issue")
		}
	}
}
