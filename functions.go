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

func main() {
	fmt.Println(sum(5, 10))
	fmt.Println(swap(20, 10))
}