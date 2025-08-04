package core

import (
	"fmt"
	"strings"
	"unicode"
)

func Evaluate(pass string) {
	pass = strings.TrimSpace(pass)

	var hasUppercase, hasLowercase, hasDigit, hasSpecial bool

	for _, r := range pass {
		switch {
		case unicode.IsDigit(r):
			hasDigit = true
		case unicode.IsLower(r):
			hasLowercase = true
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			hasSpecial = true
		case unicode.IsUpper(r):
			hasUppercase = true
		}

	}

	if len(pass) < 8 {
		fmt.Println("Password is too short (min 8 characters).")
		return
	}

	if hasUppercase && hasLowercase && hasDigit && hasSpecial {
		fmt.Println("Password is strong.")
	} else {
		fmt.Println("Password is weak.")
	}
}
