# Go Lang 

1. Initialize Module
   
    ```go
    go mod init <app-name>
    ```
2. All code must resides in package, top line would have:

    ```go
    package main
    ```
3. Entry point for compiler to start the execution from that point => zip your code in `main()` function like in c, c++
4. import the package of which you are using the pre-defined functions, like `Print()` is from **`fmt`** package
5. To run Go Program
    ```sh
    go run <file-name>.go
    ```
### Links to Navigate

-  [print](https://github.com/princebansal7/Learn-GoLang/blob/main/01.Basics/01.print.go)
-  [Variables](https://github.com/princebansal7/Learn-GoLang/blob/main/01.Basics/02.variables.go)
-  [Data types](https://github.com/princebansal7/Learn-GoLang/blob/main/01.Basics/03.dataTypes.go)
-  [User Input](https://github.com/princebansal7/Learn-GoLang/blob/main/01.Basics/04.userInput.go)
-  [Operators](https://github.com/princebansal7/Learn-GoLang/blob/main/01.Basics/05.operators.go)

### Go Advantages

- Detects errors while writing the code without even compiling
- Simple syntax like python yet efficient like C, C++
- Fast Build, startup and run time
- Resource efficient and same bianry runs on multiple OS

### About compiling

- Each module should have only 1 main function so either create seperate folders for each program or keep one main function and use functions (Like I did)
- To compile:
  ```sh
  go run .
  ```
- Pass all the files whoes functions are present in main file, so that it can compile them (compile them in any other, order of execution will be in sequence the way you called the functions)
  ```sh
  go run 00.main.go 02.variables.go 01.print.go
  ```
- Some more key points regarding compilation in **Go**
  - `go run` command is a shortcut for `go build -o temp_binary && ./temp_binary`
  - When you run `go run file1.go file2.go file3.go`
	- Go will compile all these files together, treating them as a single Go program.
	- All files must belong to the same package (typically package main) and reside in the same directory.
	- One of the files must contain a main() function.
	- All referenced functions/types must be defined across these files, or else it will throw “undefined” errors.
  - Key Rules:
    1.	Go compiles only the files you give it.
	2.	It must find a main() in the list of files.
	3.	All references across the files must be resolved within the given files. (Like if defined an function but while compiling you do not pass the reference as argument, it will give error)
    
        for example:
        
         `go run main.go`
           but didn't pass print.go but print() in present in main() => error