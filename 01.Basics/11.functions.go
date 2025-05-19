package main

import "fmt"

func functions() {

	// 1. Takes Nothing, Returns Nothing

	greet()

	// 2. Takes Something, Returns Nothing

	greetUser("Prince")

	// - Multiple parameters
	add1(10, 59)

	// - Combining types
	add2(10, 20)

	// 3. Takes Nothing, Returns Something

	number := getNumber()
	fmt.Println("Got:", number)

	// - Can also return multiple values
	x, y := getCoordinates()
	fmt.Printf("X = %v Y = %v \n", x, y)

	// 4. Takes Something, Returns Something

	result := square(5)
	fmt.Println("Square:", result)

	// - Multiple parameters, multiple returns:
	divisionResult, moduloResult := divideAndModule(10, 3)
	fmt.Printf("Division: %v, Modulo: %v\n", divisionResult, moduloResult)

	// 5. Named return value

	area := calcArea(10, 20)
	fmt.Println("Area of rectangle is:", area)

	// 6. Variadic Functions (Takes variable number of arguments)

	fmt.Println("Sum is:", sum(1, 2, 3, 4)) // => 10

	// 7. Anonymous (Inline) Functions

	func() {
		fmt.Println("Anonymous function")
	}() // note the () to invoke it

	// - with args and return
	multiplication := func(a, b int) int {
		return a * b
	}(2, 3)

	fmt.Println("Multiplication is:", multiplication) // 5

	// 8. Function Assigned to a Variable

	add := func(a, b int) int {
		return a + b
	}

	fmt.Println("Addition:", add(3, 4)) // 7

	// 9. Functions Returning Functions

	timesTwo := multiplier(2)
	fmt.Println(timesTwo(5)) // 10

	// 10. Passing Functions as Arguments
	fmt.Println("Subtraction is:", compute(minus, 3, 4)) // 7

	fmt.Println()
}

// ######### FUNCTIONS ##################

// 1.
func greet() {
	fmt.Println("Hello!")
}

// 2.
func greetUser(name string) {
	fmt.Println("Hello,", name)
}

// 2.1 Multiple parameters
func add1(a int, b int) {
	fmt.Println("Sum:", a+b)
}

// 2.2 Combining types
func add2(a, b int) { // both are int
	fmt.Println("Sum:", a+b)
}

// 3.
func getNumber() int {
	return 42
}

// 3.1 Returning multiple values
func getCoordinates() (int, int) {
	return 10, 20
}

// 4.
func square(n int) int {
	return n * n
}

// 4.1 - Multiple parameters, multiple returns (2nd pair  of () represents type of returned values):
func divideAndModule(a, b int) (int, int) {
	return a / b, a % b
}

// 5.
func calcArea(length, width int) (area int) {
	area = length * width
	return // implicit return of named value
}

// 6.
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// 9.
func multiplier(factor int) func(int) int {
	return func(val int) int {
		return val * factor
	}
}

// 10.
func compute(fn func(int, int) int, a, b int) int {
	return fn(a, b)
}

func minus(a int, b int) int {
	return a - b
}
