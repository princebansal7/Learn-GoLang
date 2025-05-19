package main

import "fmt"

func slices() {

	// - Slices are abstraction on arrays
	// - variable length,index based and can be resized when needed
	//   => provides dynamic sizing

	// 1. Declare an empty slice

	var ages []int
	fmt.Println(ages) // []

	emptySlice := make([]int, 5) // length = 5, capacity = 5, values = [0 0 0 0 0]
	fmt.Println(emptySlice)
	emptySlice2 := make([]int, 5, 10) // length = 5, capacity = 10
	fmt.Println(emptySlice2)
	emptySlice3 := make([]int, 0) // length = 0, capacity = 0
	fmt.Println(emptySlice3)      // []

	// 2. Declare slice and initialize with values

	var firstName = []string{"Ram", "John", "Pinku"}
	fmt.Println(firstName[0], firstName[1], firstName[2]) // Ram John Pinku

	firstName2 := []int{1, 2}
	fmt.Println(firstName2[0], firstName2[1]) // 1 2

	// 3. Slice from an array

	arr := [5]int{10, 20, 30, 40, 50}
	s := arr[1:4]  // elements at index 1, 2, 3 => [20 30 40] => [start:end] , end not included like python
	fmt.Println(s) // [20 30 40]

	// 5. Slice from a slice

	baseSlice := []int{1, 2, 3, 4, 5}
	subSlice := baseSlice[2:]                   // [3 4 5]
	subSlice2 := baseSlice[:3]                  // [1 2 3]
	subSlice3 := baseSlice[1:4]                 // [2 3 4]
	fmt.Println(subSlice, subSlice2, subSlice3) // [3 4 5] [1 2 3] [2 3 4]

	// 6. Appending elements to Slice [Important]

	s1 := []int{1, 2}
	s1 = append(s1, 3, 4)
	fmt.Println(s1) // [1 2 3 4]

	// Appending slices

	s2 := []int{1, 2}
	s3 := []int{3, 4}
	s4 := append(s3, s2...) // appending s2 on s3 => [3 4 1 2]
	fmt.Println(s4)         // [3 4 1 2]

	// NIL vs Empty slice

	var sl1 []int         // nil slice (len=0, cap=0, s1 == nil is true)
	sl2 := []int{}        // empty slice (len=0, cap=0, s2 == nil is false)
	fmt.Println(sl1, sl2) // [] []

	// example Appending slice elements:

	names := []string{"Prince", "John"}
	fmt.Println(names) // [Prince John]
	names = append(names, "Ram"+"-"+"Naruto", "Mike")
	fmt.Println(names) // [Prince John Ram-Naruto Mike]

	fmt.Println("Printing whole slice:", names)            // Printing whole slice: [Prince John Ram-Naruto Mike]
	fmt.Println("Printing 3rd value:", names[2])           // Printing 3rd value: Ram-Naruto
	fmt.Println("Printing length of slice:", len(names))   // Printing length of slice: 4
	fmt.Println("Printing capacity of slice:", cap(names)) // Printing capacity of slice: 4
	fmt.Printf("Printing slice type: %T \n", names)        // Printing slice type: []string
	fmt.Printf("Printing slice address: %p \n", &names)    // Printing slice address: 0x1400000c0c0
	fmt.Println("Printing address of all elements:")
	fmt.Printf("&names[0]= %v, &names[1]= %v, &names[2]= %v, &names[3]= %v\n", &names[0], &names[1], &names[2], &names[3])
	// &names[0]= 0x1400010c040, &names[1]= 0x1400010c050, &names[2]= 0x1400010c060, &names[3]= 0x1400010c070
}
