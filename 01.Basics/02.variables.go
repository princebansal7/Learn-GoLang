package main

import (
	"fmt"
)

func variables() {

	// 1. Using var with explicit type (declaration + definition)
	var a int
	a = 10
	fmt.Println("Value of a is", a)

	// 2. Using var with initialization (type inference or explicit)

	// a) With type inference:
	var b = 10 // Go infers type as int
	fmt.Println("Value of b is", b)

	// b) With explicit type:
	var c float64 = 3.14
	fmt.Println("Value of c is", c)

	// 3. Short declaration using := (inside functions only)
	//    - Go infers type from the value.
	//    - This can only be used inside functions, not at package level.
	d := "Hello"
	fmt.Println("Value of d is", d)

	// 4. Multiple variable declarations (same or different types)
	// a) Same type:
	var x, y int
	x, y = 1, 2
	fmt.Printf("Value of x=%v and y=%v \n", x, y)

	// b) With initialization:
	var X, Y = 10, "Go"
	fmt.Printf("Value of x=%v and y=%v \n", X, Y)

	// c) Short declaration:
	name, age := "Prince", 25
	fmt.Println(name, age)

	// 5. Grouped variable declaration using var block
	var (
		i int = 10
		j     = 20 // type inferred
		k string
		l = "GoLang"
	)
	fmt.Println(i, j, k, l)

	// 6. Constants (immutable values)

	const year = 1998
	const pi = 3.14
	const lang string = "Go"
	fmt.Println("Birth Year:", year)
	fmt.Println("Pi value :", pi)
	fmt.Println("string constant:", lang)

	// 7. Zero value initialization

	var p int     // 0
	var q float64 // 0.0
	var r string  // ""
	var s bool    // false
	fmt.Println(p, q, r, s)
}
