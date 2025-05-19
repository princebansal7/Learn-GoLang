package main

import (
	"fmt"
	"strings"
)

func if_else() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	// 1. if block
	if n%2 == 0 {
		fmt.Printf("%v is a Even number\n", n)
	}

	// 2. if-else block
	if n%2 == 0 {
		fmt.Printf("%v is a Even number\n", n)
	} else {
		fmt.Printf("%v is a Odd number\n", n)
	}

	// 3. else-if block

	if n > 0 {
		fmt.Printf("%v is a Positive number\n", n)
	} else if n < 0 {
		fmt.Printf("%v is a Negative number\n", n)
	} else {
		fmt.Printf("%v is neither positive nor negative\n", n)
	}

	// 4. if-else can be used input validation

	var firstName string
	var lastName string
	var email string
	fmt.Print("Enter a first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter a last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter a email: ")
	fmt.Scan(&email)

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") && len(email) >= 4

	if isValidName && isValidEmail {
		fmt.Printf("Hey %v %v! \nWelcome Aboard!\nYour email is: %v\n", firstName, lastName, email)
	} else {
		if !isValidName {
			fmt.Println("Entered first or last should be more than of length 2")
		} else if !isValidEmail {
			fmt.Println("email must be of length more than 4 and contains @")
		}
	}

	fmt.Println()
}
