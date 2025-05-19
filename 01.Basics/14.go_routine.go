package main

import (
	"fmt"
	"time"
)

func go_routine() {

	fmt.Println("Hello") // executes immediately
	welcome("Prince")    // executes after 5 seconds (we added custom delay)
	fmt.Println("Bye")   // executes after welcome() executes => by default: sequential code execution

	// => Executes on single thread (main thread)
	// - But we want the code to run on different thread in case of such delays and usually we are
	//   aware of which block for code can cause delays => like any API requests
	// => The code which can cause potential delay should be executed on separated thread, so that
	//    rest of the code does wait.

	// Handing above block using 'Goroutines'
	// - just write 'go' in front of the block and it will start a new goroutine
	// - hides all internal complexity like thread creation.
	// - A goroutine is lightweight thread managed by the 'go' runtime

	fmt.Println("Start")

	// Problem with using go keyword alone:
	//  - If the main function exits, the program terminates—even if goroutines are still running.

	go downloadFile("file1.jpg")
	go downloadFile("file2.jpg")

	fmt.Println("Doing something else...")

	// Wait for goroutines to finish
	time.Sleep(3 * time.Second)

	fmt.Println("End")

	// Problem with using go keyword alone:
	//  - If the main function exits, the program terminates—even if goroutines are still running.

	fmt.Println()
}

func welcome(name string) {

	// Sleep() function from time package stops or block the current thread (goroutine) execution for 5 seconds

	time.Sleep(5 * time.Second)
	fmt.Println("Welcome", name)
}

func downloadFile(fileName string) {
	fmt.Println("Started downloading", fileName)
	time.Sleep(2 * time.Second)
	fmt.Println("Finished downloading", fileName)
}
