package main

import (
	"fmt"
	"unicode"
)

func checkPasswordStrength(password string) string {
	var lengthOk, upper, lower, number, special bool

	if len(password) >= 12 {
		lengthOk = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			upper = true
		case unicode.IsLower(char):
			lower = true
		case unicode.IsDigit(char):
			number = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			special = true
		}
	}

	score := 0
	if lengthOk {
		score++
	}
	if upper {
		score++
	}
	if lower {
		score++
	}
	if number {
		score++
	}
	if special {
		score++
	}

	switch score {
	case 5:
		return "Strong"
	case 3, 4:
		return "Moderate"
	default:
		return "Weak"
	}
}

func main() {
	var password string
	fmt.Print("Enter password to check strength: ")
	_, err := fmt.Scanln(&password)
	if err != nil {
		fmt.Println("Error while reading password: ", err)
	}

	strength := checkPasswordStrength(password)
	fmt.Println("Password Strength:", strength)
}
