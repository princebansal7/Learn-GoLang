package main

import "fmt"

func userInput() {

	var name string
	var age int
	var year int

	// use Scan() from fmt package to take user input
	fmt.Print("Enter your name: ")
	fmt.Scan(&name) // passing memory address: same as c (to go to the address of name variable which is pointing to the memory location)

	fmt.Print("Enter age & birth year: ")
	fmt.Scan(&age, &year)
	fmt.Println("Entered name is:", name)
	fmt.Printf("Age = %v, Birth Year = %v \n", age, year)

	fmt.Println("Memory address of name:", &name)
	fmt.Println("Memory address of age:", &age)
	fmt.Println("Memory address of year:", &year)
}
