package main

import (
	"fmt"
)

// variadic function that takes a variable number of integer arguments and returns their sum
// ellipsis (...) is used to indicate that the function can accept a variable number of arguments
func summation(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}


func main() {
	result := summation(1, 2, 3, 4, 5)
	fmt.Println("First Sum:", result)

	result2 := summation(10, 20, 30)
	fmt.Println("Second Sum:", result2)
}
