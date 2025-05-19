package main

import (
	"fmt"
	"sync"
	"time"
)

// Creating WaitGroup (Step 1)
var wg = sync.WaitGroup{}

func go_routine() {

	// Example 1:

	fmt.Println("Hello") // executes immediately
	welcome("Prince")    // executes after 4 seconds (we added custom delay)
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

	// Example 2:

	// Problem with using go keyword alone:
	//  - If the main function exits, the program terminates—even if goroutines are still running.
	//  - by default the main goroutine does NOT wait for other goroutines

	fmt.Println("Hey")      // executes immediately
	go welcome("Prince!!!") // runs on other goroutine for 4 seconds as per time()
	fmt.Println("Bye Bye!") // executes immediately after "Hey", main thread completes => welcome() response doesn't show as program terminates

	// => To avoid this, we need to tell the 'main' that it needs to wait until 'welcome()' is done
	//    - to do this create 'WaitGroup': waits for launched goroutine to finish
	//    - use 'sync' package which provides basic synchronization functionality
	//    - WaitGroup provides 3 functions:
	//      - Add(): Sets the number of goroutines to wait for (add before goroutine functions, it increases the counter by provided number)
	//      - Wait(): Blocks until the WaitGroup counter is 0 (usually put at the end of main)
	//      - Done(): Decrements the WaitGroup counter by 1, called by the goroutine to indicate that it's finished (used in the function for which
	//                goroutine is created for)

	fmt.Println("Hey Hey")

	wg.Add(1) // (Step 2) adds below goroutine to WaitGroup (and increases the counter by 1, showing 1 goroutine is added in WaitGroup)
	// Note: if we had 1 more goroutine function: we would do: wg.Add(2), indicating WaitGroup have 2 goroutines

	go welcome2("Prince Bansal")
	fmt.Println("BBye!")

	fmt.Println()
	wg.Wait() // (Step 3) tells main thread to wait, until all goroutine are completed (when counter gets to 0)
}

func welcome(name string) {
	// Sleep() function from time package stops or block the current thread (goroutine) execution for 5 seconds
	time.Sleep(4 * time.Second)
	fmt.Println("Welcome", name)
}

func welcome2(name string) {
	// Sleep() function from time package stops or block the current thread (goroutine) execution for 5 seconds
	time.Sleep(5 * time.Second)
	fmt.Println("Welcome", name)
	wg.Done() // (Step 4) Tells that this goroutine is done executing, and decrements the WaitGroup counter by 1
}

/*
OUTPUT:

	Hello
	Welcome Prince
	Bye
	Hey
	Bye Bye!
	Hey Hey
	BBye!

	Welcome Prince!!!
	Welcome Prince Bansal

*/
