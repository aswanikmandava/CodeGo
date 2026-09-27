package main

import (
	"fmt"
	"errors"
)

// function that uses a named return value to return the sum of two integers
func add(a int, b int) (sum int) {
	// named return value 'sum' is declared in the function signature
	sum = a + b
	return // returning the named return value
}

// function that takes 2 integers as arguments and return 2 names parameters
// if both arguments are of same type, specify type for the last one
func divide(a, b int) (result int, err error) {
	if b == 0 {
		err = errors.New("division by zero") // setting the error return value
		return // returning the named return values (result and err)
	}
	result = a / b
	return // returning the named return values (result and err)
}

func main() {
	result := add(5, 10)
	fmt.Println("Sum:", result)

	result2, error := divide(10, 0)
	fmt.Println("result: ", result2, " and error: ", error)

	result3, err := divide(100, 50)
	fmt.Println("result3: ", result3, " and error: ", err)
}
