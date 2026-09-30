package main

import (
	"fmt"
)

// defining a struct
type Person struct {
	// default value of name is empty string
	name string
	// default value of age is 0
	age  int
	// default value of employed is false
	employed bool
}

func main() {
	// creating an instance of the struct
	p1 := Person{name: "John", age: 30, employed: true}
	p2 := Person{name: "Alice", age: 25, employed: false}
	p3 := Person{name: "Bob"}
	p4 := Person{}
	// assigning values to the fields of the struct
	p3.age = 35
	p4.name = "Eve"
	p4.age = 28
	p4.employed = false
	fmt.Println("p1: ", p1)
	fmt.Println("p2: ", p2)
	fmt.Println("p3: ", p3)
	fmt.Println("p4: ", p4)
}