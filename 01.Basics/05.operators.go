package main

import "fmt"

func operators() {

	// Arithmetic Operators: +, -, /, *, %
	// Relational Operators: <, >, <=, >=, !=, ==
	// Logical Operators: &&, ||, !
	// Bitwise Operators: &, |, ^, <<, >>, New operator: &^ (AND NOT) is a bit clear operator
	// Assignment Operators: =, +=, -=, /=, *=, %=, &=, |=, ^=, <<=, >>=
	// Misc Operators: & (Address of: returns address), * (for pointer variable), <- (receive operator: receives a value from channel)

	var value = 10
	value++
	fmt.Println(value) // 11
	value--
	fmt.Println(value) // 10
	// --value // Error
	// ++value // Error
}
