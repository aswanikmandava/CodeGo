package main

import (
	"fmt"
)

// function that takes a slice of integers and updates the first item
// slice is passed by reference, so changes made inside the function will affect the original slice
func updateSliceFirstItem(s []int) {
	s[0] = 100 // updating the first item of the slice
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Before:", numbers)
	updateSliceFirstItem(numbers)
	fmt.Println("After:", numbers)
}
