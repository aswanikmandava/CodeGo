package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var b int = 20
	var c int = 30
	// store the address of c in d
	var d *int = &c
	// modify the value of c using d
	// dereference d and assign a new value to c
	*d = 40
	fmt.Println("Value of c:", c)
	fmt.Println("---------Arithmetic-----------")
	fmt.Printf("addition of %d and %d: %d", a, b, a+b)
	fmt.Printf("\nsubtraction of %d and %d: %d", a, b, a-b)
	fmt.Printf("\nmultiplication of %d and %d: %d", a, b, a*b)
	fmt.Printf("\ndivision of %d and %d: %d", a, b, a/b)
	fmt.Printf("\nmodulus of %d and %d: %d", a, b, a%b)
	fmt.Println("--------- Conditional ---------")
	fmt.Printf("%d != %d: %t", a, b, a!=b)
	fmt.Printf("\n%d == %d: %t", a, b, a==b)
	fmt.Printf("\n%d <= %d: %t", a, b, a<=b)
	fmt.Printf("\n%d >= %d: %t", a, b, a>=b)
	fmt.Println("--------- Logical ---------")
	fmt.Printf("%t && %t: %t", a!=b, a==b, a!=b && a==b)
	fmt.Printf("\n%d || %d: %t", a, b, a==b || a!=b)
	fmt.Printf("\n!%t: %t", a==b, !(a==b))
}
