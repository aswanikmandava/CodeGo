package main

import "fmt"

func main() {
	name := "Aswani"

	namePtr := &name

	fmt.Println(&namePtr)

	printPtr(namePtr)
}

func printPtr(namePointer *string) {
	fmt.Println(&namePointer)
}