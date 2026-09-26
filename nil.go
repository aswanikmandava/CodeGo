package main

import (
	"fmt"
)

func main() {
	var x *int
	var mySlice []string
	var myMap map[string]int
	var y int

	// nil is used to represent types that have no value, such as pointers, slices, maps, channels, and interfaces
	if x == nil {
		fmt.Println("x is nil")
	}
	if mySlice == nil {
		fmt.Println("mySlice is nil")
	}
	if myMap == nil {
		fmt.Println("myMap is nil")
	}
	// following line will not compile because y is an int and cannot be compared to nil
	// if y == nil {
	// 	fmt.Println("y is nil")
	// }
}
