package main

import "fmt"

func main() {
	var i, j int = 1, 2 // multiple variable declaration with explicit types
	k, l := 3, 4 // multiple variable declaration with implicit types using short variable declaration

	var m int // m is an integer variable with default value 0
	var n float64 // n is a float64 variable with default value 0
	var o bool	// o is a boolean variable with default value false
	var p string // p is a string variable with default value ""
	fmt.Println(i, j, k, l) // prints: 1 2 3 4
	fmt.Println(m, n, o, p) // prints default values of variables: 0 0 false ""
}
