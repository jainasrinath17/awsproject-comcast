package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Function to check for 4 consecutive repeating digits
func hasConsecutiveRepeats(card string) bool {
	for i := 0; i < len(card)-3; i++ {
		if card[i] == card[i+1] && card[i+1] == card[i+2] && card[i+2] == card[i+3] {
			return true
		}
	}
	return false
}

func isValidCard(card string) bool {
	// Checking for invalid spaces and hyphen combinations
	if strings.Contains(card, " ") {
		return false
	}

	// Defining the regex for valid cards
	// Ensuring that the card starts with 4, 5, or 6, followed by 15 digits
	// Allowing hyphens in the form of xxxx-xxxx-xxxx-xxxx but no spaces or other separators
	cardRegex := regexp.MustCompile(`^[456]\d{3}-?\d{4}-?\d{4}-?\d{4}$`)

	// Checking if card matches the valid format
	if !cardRegex.MatchString(card) {
		return false
	}

	// Removing hyphens for further validation
	card = strings.ReplaceAll(card, "-", "")

	// Checking if card has 16 digits after removing hyphens
	if len(card) != 16 {
		return false
	}

	// Checking for repeated consecutive digits
	if hasConsecutiveRepeats(card) {
		return false
	}

	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// number of input lines
	var n int
	fmt.Scan(&n)

	// Processing each card number
	for i := 0; i < n; i++ {
		scanner.Scan()
		card := scanner.Text()

		if isValidCard(card) {
			fmt.Println("Valid")
		} else {
			fmt.Println("Invalid")
		}
	}
}
