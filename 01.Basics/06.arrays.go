package main

import "fmt"

func arrays() {

	// 1. Declare an array without initializing values

	var ages [10]int     // fixed size: 10 and default value in int type array = 0
	fmt.Println(ages[2]) // 0

	ages[0] = 23
	ages[2] = 10
	ages[1] = 27
	fmt.Println(ages[0], ages[1], ages[2], ages[5]) // 23 27 10 0

	// 2. Declare and initialize with values

	var firstName = [5]string{"Ram", "John", "Pinku"}
	fmt.Println(firstName[0], firstName[1], firstName[2], firstName[4]) // Ram John Pinku ""

	// 3. Use ... to let Go infer the length

	arr := [...]int{1, 2, 3}
	fmt.Println(arr[1]) // 2

	// 4. Partial initialization

	arr2 := [5]int{1: 10, 3: 69}

	//    - Creates [0 100 0 200 0]
	//    - Only index 1 and 3 are explicitly set.

	fmt.Println(arr2[0], arr2[1], arr2[2], arr2[3], arr2[4]) // 0 10 0 69 0

	fmt.Println("Printing whole array:", arr2)          // Printing whole array: [0 10 0 69 0]
	fmt.Println("Printing 2nd value:", arr2[1])         // Printing 2nd value: 10
	fmt.Println("Printing length of array:", len(arr2)) // Printing length of array: 5
	fmt.Printf("Printing array type: %T \n", arr2)      // Printing array type: [5]int
	fmt.Printf("Printing array address: %p \n", &arr2)  // printing array address:  0x1400012e000 (same as the first element)
	fmt.Println("Printing address of all elements:")
	fmt.Printf("&arr2[0]= %v, &arr2[1]= %v, &arr2[2]= %v, &arr2[3]= %v, &arr2[4]=% v \n", &arr2[0], &arr2[1], &arr2[2], &arr2[3], &arr2[4])
	// &arr2[0]= 0x1400012e000, &arr2[1]= 0x1400012e008, &arr2[2]= 0x1400012e010, &arr2[3]= 0x1400012e018, &arr2[4]= 0x1400012e020

	// 5.  Declaring a multidimensional arrays

	var matrix1 [2][3]int = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(matrix1[0], matrix1[0][0], matrix1[1][2]) // [1 2 3] 1 6

	matrix2 := [...][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(matrix2[0][2], matrix2[1], matrix2[1][1]) // 3 [4 5 6] 5

}
