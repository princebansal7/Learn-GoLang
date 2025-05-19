package main

import "fmt"

func structType() {
	// A struct in Go is a composite data type (like a class in other languages, but without methods or inheritance).
	// It groups related data together into one unit.
	// - to collect different data types and make single type from them

	// LIMITATIONS:
	// No inheritance (use composition)
	// No constructors (but you can create factory functions)
	// Cannot declare methods inside struct blocks

	// Creating custom type using 'struct'
	// - 'type' keyword creates a new type with the name you specify, in this example: Person
	//    => Create a named 'Person'based on a structure of Name, Age, .... etc

	type Person struct {
		Name string
		Age  int
	}

	// 1. Declaring and Initializing Structs

	// a) Named Fields
	p1 := Person{
		Name: "Prince",
		Age:  27,
	}
	fmt.Println(p1)

	// b) Positional Values (Order matters)
	// Avoid this unless all fields are obvious and stable.
	p2 := Person{"Prince", 27}
	fmt.Println(p2)

	// c) Zero-value Initialization
	var p3 Person // p.Name == "", p.Age == 0
	fmt.Println(p3)

	// d) Using new keyword
	p4 := new(Person)
	p4.Name = "Prince"
	p4.Age = 25
	fmt.Println(p4.Name, p4.Age)

	// Accessing and updating values
	p4.Name = "Prince Bansal"
	p4.Age = 27
	fmt.Println(p4.Name, p4.Age)

	// 2. Anonymous Structs
	// Useful for temporary or quick use.
	person := struct {
		Name string
		Age  int
	}{
		Name: "Prince",
		Age:  25,
	}
	fmt.Println(person)

	// 3. Anonymous Fields (Embedding)

	type Address struct {
		City string
	}
	type Employee struct {
		Name    string
		Address // anonymous field
	}

	e := Employee{Name: "Prince", Address: Address{City: "Chandigarh"}}
	fmt.Println(e.City) // Access City directly
	fmt.Println(e.Name)

	// 4. Struct Slices and Arrays
	people := []Person{
		{Name: "Ram", Age: 20},
		{Name: "John", Age: 30},
	}
	fmt.Println(people)

	fmt.Println()
}
