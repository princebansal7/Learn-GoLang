package main

import "fmt"

func loops() {
	// - GO have only one type of loop - for loop

	// 1. classic for loop (like c/java)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 2. While-style loop (using only condition)

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// 3. Infinite Loop

	for {
		fmt.Println("Infinite loop")
		// Use break or some condition to exit
		break
	}

	// 4. Loop over a slice or array (using range)

	nums := []int{10, 20, 30, 40, 50}
	for index, value := range nums {
		fmt.Println("Index:", index, ", Value:", value)
	}
	// ignore index
	for _, value := range nums {
		fmt.Println("Value:", value)
	}

	// ignore value
	for index := range nums {
		fmt.Println("Index:", index)
	}

	// 5. Loop over a map

	m := map[string]int{"a": 1, "b": 2}
	for key, value := range m {
		fmt.Println(key, value)
	}

	// 6. Loop over a string
	s := "hello"
	for i, ch := range s {
		fmt.Printf("%d -> %c\n", i, ch)
	}

	// 7. Breaking out of loop (break)
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Print(i, " ") // prints till 4 then gets out of the loop
	}
	fmt.Println()

	// 8. Skip current iteration (continue)

	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Print(i, " ") // prints only even numbers
	}

	fmt.Println()

	// 9. Loop with labels (useful for nested loops)

outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j > 4 {
				break outer
			}
			fmt.Println(i, j)
		}
	}

	fmt.Println()
}
