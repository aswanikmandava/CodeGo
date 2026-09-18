package main

import "fmt"

func main() {
	defer fmt.Println("This will be printed last after the main function completes.")
	for i := 0; i < 3; i++ {
		defer fmt.Println("Deferred call number:", i)
	}
	fmt.Println("Hello, World!")
}
