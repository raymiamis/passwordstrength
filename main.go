package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"
	"unicode"
)

func isPwned(password string) (bool, int, error) {

	hasher := sha1.New()
	hasher.Write([]byte(password))
	hash := strings.ToUpper(hex.EncodeToString(hasher.Sum(nil)))

	prefix := hash[:5]
	suffix := hash[5:]

	url := "https://api.pwnedpasswords.com/range/" + prefix
	resp, err := http.Get(url)
	if err != nil {
		return false, 0, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return false, 0, fmt.Errorf("API-Error: %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, 0, err
	}

	lines := strings.Split(string(body), "\n")
	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}
		if strings.TrimSpace(parts[0]) == suffix {

			count := strings.TrimSpace(parts[1])
			var n int
			_, err := fmt.Sscanf(count, "%d", &n)
			if err != nil {
				return false, 0, err
			}
			return true, n, nil
		}
	}

	return false, 0, nil
}

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
	fmt.Println(`
 ▄▄▄· ▄▄▄· .▄▄ · .▄▄ · ▄▄▌ ▐ ▄▌      ▄▄▄  ·▄▄▄▄  .▄▄ · ▄▄▄▄▄▄▄▄  ▄▄▄ . ▐ ▄  ▄▄ • ▄▄▄▄▄ ▄ .▄
▐█ ▄█▐█ ▀█ ▐█ ▀. ▐█ ▀. ██· █▌▐█▪     ▀▄ █·██▪ ██ ▐█ ▀. •██  ▀▄ █·▀▄.▀·•█▌▐█▐█ ▀ ▪•██  ██▪▐█
 ██▀·▄█▀▀█ ▄▀▀▀█▄▄▀▀▀█▄██▪▐█▐▐▌ ▄█▀▄ ▐▀▀▄ ▐█· ▐█▌▄▀▀▀█▄ ▐█.▪▐▀▀▄ ▐▀▀▪▄▐█▐▐▌▄█ ▀█▄ ▐█.▪██▀▐█
▐█▪·•▐█ ▪▐▌▐█▄▪▐█▐█▄▪▐█▐█▌██▐█▌▐█▌.▐▌▐█•█▌██. ██ ▐█▄▪▐█ ▐█▌·▐█•█▌▐█▄▄▌██▐█▌▐█▄▪▐█ ▐█▌·██▌▐▀
.▀    ▀  ▀  ▀▀▀▀  ▀▀▀▀  ▀▀▀▀ ▀▪ ▀█▄▀▪.▀  ▀▀▀▀▀▀•  ▀▀▀▀  ▀▀▀ .▀  ▀ ▀▀▀ ▀▀ █▪·▀▀▀▀  ▀▀▀ ▀▀▀ ·`)
	fmt.Print("\nEnter password to check strength: ")
	_, err := fmt.Scanln(&password)
	if err != nil {
		fmt.Println("Error while reading password: ", err)
	}

	strength, feedback := checkPasswordStrength(password)
	fmt.Println("Password Strength:", strength)

	if len(feedback) > 0 {
		fmt.Println("------------------------------------------------------------------------------------")
		fmt.Println("Feedback:")
		for _, msg := range feedback {
			fmt.Println(msg)
		}
	}

	fmt.Println("------------------------------------------------------------------------------------")
	pwned, count, err := isPwned(password)
	if err != nil {
		fmt.Println("Error while checking pwned status:", err)
	} else if pwned {
		fmt.Printf("Warning! This password has been compromised %d times (source: HaveIBeenPwned).\n", count)
	} else {
		fmt.Println("Nice! This password hasn't been found in any leaks (source: HaveIBeenPwned).")
	}
}
