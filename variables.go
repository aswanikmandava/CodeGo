package main

import "fmt"

var a int = 10 // a is an integer variable with value 10
var b = 20 // b is an integer variable with value 20, type inferred
var c int // c is an integer variable with default value 0
const d = 40 // d is a constant with value 40 which cannot be changed

func main() {
	var a int = 30 // variable a in function scope, shadows package-level variable a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
