package main

import "fmt"

// type assertion
func assertType(v any) {
	if s, ok := v.(string); ok {
		fmt.Println("Got a string", s)
	} else if i, ok := v.(int); ok {
		fmt.Println("Got an int", i)
	} else if f, ok := v.(float64); ok {
		fmt.Println("Got a float", f)
	} else if b, ok := v.(bool); ok {
		fmt.Println("Got a boolean", b)
	} else {
		fmt.Println("Unknown")
	}
}

func main() {
	assertType("Hello")
	assertType(10)
	assertType(5.12)
	assertType(true)
}
