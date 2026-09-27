package main

import (
	"fmt"
	"time"
) 
	

func greet() {
	fmt.Println("Hello from Goroutine")
}

func main() {

	// invoke function call as a goroutine
	// syntax: go <func_name>()

	go greet()

	// invoke anonymous function as a goroutine
	go func(m string) {
		fmt.Println(m)
	}("hello from anonymous go")

	// sleep to allow goroutine to complete before the main program
	// since main program and goroutine runs concurrently
	time.Sleep(100 * time.Millisecond)
	fmt.Println("End of program")

}
