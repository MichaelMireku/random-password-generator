package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	fmt.Println("Random Password Generator")
	fmt.Println("--------------------------")

	var passwordLength int
	var includeSpecialChars bool

	// Get user input for password length
	fmt.Print("Enter the desired password length: ")
	fmt.Scan(&passwordLength)

	// Get user input for including special characters
	fmt.Print("Include special characters? (y/n): ")
	var specialCharInput string
	fmt.Scan(&specialCharInput)
	includeSpecialChars = strings.ToLower(specialCharInput) == "y"

	// Generate a random password
	password := generateRandomPassword(passwordLength, includeSpecialChars)

	// Print the generated password
	fmt.Println("\nYour random password is:", password)

	// Wait for user input before exiting
	fmt.Print("\nPress Enter to exit...")
	fmt.Scanln()
	clearScreen()
}

func generateRandomPassword(length int, includeSpecialChars bool) string {
	// Define character sets
	var characters string
	alphaLower := "abcdefghijklmnopqrstuvwxyz"
	alphaUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	specialChars := "!@#$%^&*()-=_+"

	// Include character sets based on user preferences
	characters += alphaLower + alphaUpper + digits
	if includeSpecialChars {
		characters += specialChars
	}

	// Validate password length
	if length <= 0 {
		fmt.Println("Invalid password length. Defaulting to length 12.")
		length = 12
	}

	// Initialize an empty password string
	password := ""

	// Set the seed for random number generation
	rand.Seed(time.Now().UnixNano())

	// Generate the password by randomly selecting characters
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(characters))
		password += string(characters[randomIndex])
	}

	return password
}

func clearScreen() {
	cmd := exec.Command("clear") // For Windows use "cls"
	cmd.Stdout = os.Stdout
	cmd.Run()
}
