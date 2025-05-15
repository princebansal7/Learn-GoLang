package main

import "fmt"

func print() {
	fmt.Print("Hello World!")

	fmt.Print("\n")
	//    or
	fmt.Println()

	// formatted output
	var name = "Prince"
	var gender = "Male"
	fmt.Printf("hello %v \n, Gender: %v", name, gender) // %v = variable reference in default format
}
