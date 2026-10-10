package main

import (
	"fmt"
)

func main() {
	// declaring a pointer variable of type int
	var ptr *int
	// & operator is used to get the address of a variable
	var num int = 42
	ptr = &num
	// * operator is used to dereference a pointer and access the value it points to
	fmt.Println("Value of num:", *ptr)
}
