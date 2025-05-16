package main

import (
	"fmt"
)

func variables() {
	var a = 5
	var b = 10
	fmt.Println("Value of a is", a, "and b is", b)

	var firstName = "Prince"
	var lastName = "Bansal"
	fmt.Println("First Name:", firstName, "\nLast Name:", lastName)

	// constant values

	const year = 1998
	fmt.Println("Birth Year:", year)
}
