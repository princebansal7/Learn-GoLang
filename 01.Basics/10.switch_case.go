package main

import "fmt"

func switch_case() {

	// 1. classic switch-case (doesn't need break after each case, once matches automatically breaks)
	//   - Expression after switch can be of any comparable type: int, string, bool, etc.
	//   - Cases must be of the same type and comparable to the switch expression.

	x := 2
	switch x {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other")
	}

	// 2. Multiple values in a case

	x = 1
	switch x {
	case 1, 3, 5: // 1 or 3 or 5 => matches this case
		fmt.Println("Odd")
	case 2, 4, 6:
		fmt.Println("Even")
	}

	// 3. Switch without an expression (acts like if-else)
	x = 10
	switch {
	case x < 0:
		fmt.Println("Negative")
	case x == 0:
		fmt.Println("Zero")
	default:
		fmt.Println("Positive")
	}

	// 5. Fallthrough keyword
	// - Go does not fall through by default (unlike C/C++). You have to explicitly use fallthrough if needed:

	x = 1
	switch x {
	case 1:
		fmt.Println("One")
		fallthrough // matches and still goes to the next case using 'fallthrough' keyword
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	// 6. Initialization Statement in switch
	//    - initialized variable scope is limited to switch-case block in this case

	switch x := 5 * 2; x {
	case 10:
		fmt.Println("Ten Ten ten")
	}

	// 7. Switch on Strings, Booleans, Struct Fields, Constants
	//    (Below all valid)

	// switch name {
	// case "Alice", "Bob":
	// 	fmt.Println("Known")
	// }

	// switch true {
	// case age > 18:
	// 	fmt.Println("Adult")
	// }

	// switch user.role {
	// case "admin":
	// 	fmt.Println("Admin user")
	// }

	fmt.Println()
}
