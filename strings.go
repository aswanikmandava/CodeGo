package main

import (
	"fmt"
)

func main() {

	var str string = "Hello, World!"
	// string is immutable in Go, meaning its value cannot be changed after it is created
	str = "Hello, Go!" // this creates a new string and assigns it to str
	// str[0] = 'h' // this will cause a compile-time error because strings are immutable
	// string length
	fmt.Println("Length of str:", len(str))
	greeting := "Hello" + ", " + "aswani"
	fmt.Println("Greeting:", greeting)

	// extracting a substring from a string
	substr := str[7:12] // extracting substring from index 7 to 12
	fmt.Println("Substring:", substr)

	// printing each character of the string
	for i := 0; i < len(str); i++ {
		fmt.Printf("Character at index %d: %c\n", i, str[i])
	}

	// defining a multiline string using backticks
	multilineStr := `This is a
multiline string
in Go.`
	fmt.Println("Multiline String:\n", multilineStr)

	// reference to a string variable in a multiline string
	name := "Alice"
	multilineStrWithVar := fmt.Sprintf(`Hello, %s!
Welcome to Go programming.`, name)
	fmt.Println("Multiline String with Variable:\n", multilineStrWithVar)
}
