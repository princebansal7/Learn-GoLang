package main

import "fmt"

func if_else() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	// if block
	if n%2 == 0 {
		fmt.Printf("%v is a Even number\n", n)
	}

	// if-else block
	if n%2 == 0 {
		fmt.Printf("%v is a Even number\n", n)
	} else {
		fmt.Printf("%v is a Odd number\n", n)
	}

	// else-if block

	if n > 0 {
		fmt.Printf("%v is a Positive number\n", n)
	} else if n < 0 {
		fmt.Printf("%v is a Negative number\n", n)
	} else {
		fmt.Printf("%v is neither postive nor negative\n", n)
	}

	fmt.Println()
}
