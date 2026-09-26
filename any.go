package main

import (
	"fmt"
)

// function that takes any type value and returns its type as a string
func getType(val any) string {
	return fmt.Sprintf("%T", val)
}

func main() {
	fmt.Println(getType(42))
	fmt.Println(getType("hello"))
	fmt.Println(getType(true))
}
