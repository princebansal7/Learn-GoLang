package main

import "fmt"

func maps() {
	// - Maps unique keys to values
	// - can you retrieve the most
	// - All 'keys' have same data type
	// - All 'values' have same data type

	// Zero value => A declared but uninitialized map is nil. You must use make() or a literal to initialize it.
	// Order => Maps in Go are not ordered. Don’t rely on iteration order.
	// Keys => Keys must be of a comparable type (e.g., int, string, structs with comparable fields).
	// Concurrent access => Maps are not safe for concurrent write access. Use sync.Map or mutexes for thread safety.

	// Syntax: map[KeyType]ValueType

	// 1. using make()
	m1 := make(map[string]int)
	fmt.Println(m1)

	// 2. Using map literals
	m2 := map[string]int{
		"apple":  5,
		"banana": 3,
	}
	fmt.Println(m2)

	// 3. Declare then assign
	var m3 map[string]int
	m3 = make(map[string]int)

	// Note: Without make, an uninitialized map (var m map[string]int) is nil and cannot be written to.

	// Accessing values
	m3["a"] = 10
	m3["b"] = 20
	fmt.Println(m3)
	fmt.Println(m3["a"])
	fmt.Println(m3["b"])

	// Check if Key Exists
	val, exists := m2["apple"]
	if exists {
		fmt.Println("Found:", val)
	} else {
		fmt.Println("Not found")
	}

	// Delete a Key
	m4 := map[string]int{
		"apple":   5,
		"banana":  3,
		"papaya":  2,
		"avocado": 10,
		"guava":   12,
	}
	fmt.Println("Before deleting: ", m4)
	fmt.Println("Length of map:", len(m4))
	delete(m4, "banana")
	fmt.Println("After deleting:", m4)
	fmt.Println("Length of map:", len(m4))

	// looping over a map
	for key, value := range m4 {
		fmt.Printf("%s: %d\n", key, value)
	}

	fmt.Println()
}
