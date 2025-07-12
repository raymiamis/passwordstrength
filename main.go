package main

import (
	"fmt"
	"unicode"
)

func checkPasswordStrength(password string) (string, []string) {
	var lengthOk, upper, lower, number, special bool
	var feedback []string

	if len(password) >= 12 {
		lengthOk = true
	} else {
		feedback = append(feedback, "Password shorter than 12 characters.")
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

	if !upper {
		feedback = append(feedback, "Capital letters missing.")
	}
	if !lower {
		feedback = append(feedback, "Lowercase letters missing.")
	}
	if !number {
		feedback = append(feedback, "Numbers missing.")
	}
	if !special {
		feedback = append(feedback, "Special characters missing.")
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

	var strength string
	switch score {
	case 5:
		strength = "Strong"
	case 3, 4:
		strength = "Moderate"
	default:
		strength = "Weak"
	}

	return strength, feedback
}

func main() {
	var password string
	fmt.Print("Enter password to check strength: ")
	_, err := fmt.Scanln(&password)
	if err != nil {
		fmt.Println("Error while reading password: ", err)
	}

	strength, feedback := checkPasswordStrength(password)
	fmt.Println("Password Strength:", strength)

	if len(feedback) > 0 {
		fmt.Println("Feedback:")
		for _, msg := range feedback {
			fmt.Println(msg)
		}
	}
}
