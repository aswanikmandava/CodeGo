package main

import "fmt"
// package initialization follows an order
//	1. Package-level variables are initialized
//	2. The package's init() runs
//	3. The program main() runs
//
// If multiple packages are imported, Go initializes them before main() executes
// This ensures each package is ready to use before the program starts running

var myMap = map[string]string{}

func registerMyMap(k string, v string) {
	myMap[k] = v
}

func init() {
	registerMyMap("in", "India")
}


func main() {
	fmt.Println(myMap)
}
