package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is a simple Go program demonstrating the use of defer.")
	defer fmt.Println("This message will be printed last.")
	fmt.Println("Last line of program.")

	for i := 0; i < 3; i++ {
		// Using defer inside a loop to demonstrate that deferred calls are executed in LIFO order
		defer fmt.Printf("Deferred message %d\n", i)
	}

}
