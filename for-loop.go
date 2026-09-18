package main

import "fmt"

func main() {
	var sum int // variable declaration with default value 0
	for i := 0; i < 5; i++ {
		sum += i
		fmt.Println("Current value of i:", i, "Current sum:", sum)
	}
	fmt.Println("Result: ", sum)
}
