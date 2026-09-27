package main

import "fmt"

// sum() takes two integers and returns single integer
func sum(a int, b int) int {
	return a + b
}

// swap() takes two integers and returns two integers
func swap(a int, b int) (int, int) {
	return b, a
}

// multiply() takes a pointer to an integer and multiplies its value by 2
func multiply(a *int) {
	*a *= 2
}

// function returns multiple values
func getValues() (int, string) {
	return 10, "Hello"
}

func main() {
	// arguments passed to the function by value
	fmt.Println(sum(5, 10))
	fmt.Println(swap(20, 10))
	
	// argument passed to the function by reference
	var c int = 15
	multiply(&c)
	fmt.Println("Value of c after multiplication:", c)
	var a, b = getValues()
	fmt.Printf("getValues() returned %d and %s", a, b)
	fmt.Println()
}