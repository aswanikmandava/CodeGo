package main

import "fmt"

func main() {

	// using interface{} allows values of different types
	// Starting with Go 1.18, Go introduced any as a built-in alias for interface{}.
	m := map[string]interface{}{
		"name": "Kevin",
		"age": 42,
		"isAsian": true,
	}

	// following does the same
	// m := map[string]any{
	// 	"name": "Kevin",
	// 	"age": 42,
	// 	"isAsian": true,
	// }

	// change key value if it exists
	if value, ok := m["age"]; ok {
		fmt.Println("current age is ", value)
		m["age"] = 43
	} else {
		fmt.Println("age key does not exist")
	}

	fmt.Println("modified age is ", m["age"])

	// type conversion (casting) - you use this when you want to change a value from one concrete type to another compatible concrete type
	// you are telling the compiler to structurally change the data
	var x float64 = 42.5
    var y int = int(x)          // Type Conversion: Converts 42.5 to 42
	fmt.Println("y = ", y)


	// type assertion
	// In Go (Golang), the syntax m[key].(int) is used for a type assertion on a value retrieved from a map.
    // It is how you tell the Go compiler: "I know this value is stored as a generic interface, but I want to treat it specifically as an integer."
	// you use "type assertion" when a value is hidden behind an interface (interface{} or any) and you want to reveal its true identity
	// you are not changing the data; you are asserting what it already is 

	//  SAFE: Checks if the value is actually an int
	if value, ok := m["port"].(int); ok {
		// The type assertion succeeded! 
		// 'value' is now a normal 'int' you can use here.
		fmt.Println("The port is:", value)
	} else {
		// The assertion failed (either the key doesn't exist, or it's not an int).
		fmt.Println("Error: 'port' is missing or not a valid integer.")
	}

	if value, ok := m["isAsian"].(bool); ok {
		fmt.Print("Assertion succeeded. isAsian - ", value)
	} else {
		fmt.Println("Error: 'isAsian' is missing or not a valid boolean.")
	}

}