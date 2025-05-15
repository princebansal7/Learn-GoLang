package main

import "fmt"

func dataTypes() {
	// - declaring variable without type (will give error, as expects type)
	// - It Makes code more robust and less error prone as helps devs to catch errors at compile time
	// var value

	// assigning value to declared variable
	// value = "Hello"

	// - But when we define variable like:
	// - Go does type inference => detects type itself

	// var name = "Prince"
	// var age = 27

	// - Giving type while declaring

	var name string
	var age int

	name = "Prince"
	age = 27

	fmt.Printf("Name: %v, Age: %v \n", name, age)

	fmt.Println("Printing Types:")

	fmt.Printf("type(name) = %T \ntype(age) = %T \n", name, age)
	fmt.Printf("type(69) = %T \ntype(5.66) = %T \ntype(\"Prince\") = %T \ntype(true) = %T \n", 69, "Prince", 5.66, true)
	// Output:
	// Name: Prince, Age: 27
	// Printing Types:
	// type(name) = string
	// type(age) = int
	// type(69) = int
	// type(5.66) = string
	// type("Prince") = float64
	// type(true) = bool

	// Syntactic sugar: Other way to define variables

	var name2 string = "Prince"

	// We can use below, it will do the same

	name3 := "Prince"

	// - Above works only for variables and not for constants
	// - doesn't work when we explicitly defines a type

	// const year := 20  // Error
	// year uint := 1998 // Error

	println(name2, name3)

}
