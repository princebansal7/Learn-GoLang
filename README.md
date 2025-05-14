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