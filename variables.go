package main

import "fmt"

// global variables
// package-level variable declarations
var a int = 10 // a is an integer variable with value 10
var b = 20 // b is an integer variable with value 20, type inferred
var c int // c is an integer variable with default value 0
const d = 40 // d is a constant with value 40 which cannot be changed

func main() {
	// local variable declarations
	// variable a declared in function scope, shadows package-level variable a
	var a int = 30 // variable a in function scope, shadows package-level variable a

	// shorthand variable declaration in function scope
	b := 50 // b is a new variable in function scope, shadows package-level variable b
	fmt.Printf("a = %d\n", a)
	fmt.Printf("b = %d\n", b)
	fmt.Printf("c = %d\n", c)

	// overwriting the value of b
	b = 60 // b is now 60 in function scope
	fmt.Printf("b = %d\n", b)
}
