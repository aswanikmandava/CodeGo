package main

import "fmt"

// panic stops the program due to invalid state or other abnormal condition

func main() {

	fmt.Println("start program")
	panic("Something went wrong")
	fmt.Println("this will not run")

}