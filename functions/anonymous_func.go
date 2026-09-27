package main

import (
	"fmt"
)

func main() {
	// inline function with no name is assigned to a variable
	s := func() {
		fmt.Println("Say Hi!")
	}
	// calling the anonymous function using s which has its reference
	s()

	p := func(a, b int) int {
		return a + b
	}
	fmt.Println(p(5, 10))
}
