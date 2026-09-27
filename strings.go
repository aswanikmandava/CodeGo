package main

import (
	"fmt"
	"strings"
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
	substr := str[7:9] // extracting substring from index 7 to 9
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

	// check if a string contains a substring
	myString := "Hello, Go! Welcome to the world of Go programming."
	containsGo := strings.Contains(myString, "Go")
	fmt.Println("String contains 'Go':", containsGo)

	// check if a string starts with a specific prefix
	startsWithHello := strings.HasPrefix(myString, "Hello")
	fmt.Println("String starts with 'Hello':", startsWithHello)

	// check if a string ends with a specific suffix
	endsWithProgramming := strings.HasSuffix(myString, "programming.")
	fmt.Println("String ends with 'programming.':", endsWithProgramming)

	// convert a string to uppercase
	upperStr := strings.ToUpper(myString)
	fmt.Println("Uppercase String:", upperStr)

	// find the index of a substring in a string
	index := strings.Index(myString, "Go")
	fmt.Println("Index of 'Go':", index)

	// find the last index of a substring in a string
	lastIndex := strings.LastIndex(myString, "Go")
	fmt.Println("Last Index of 'Go':", lastIndex)

	// split a string into a slice of substrings based on a delimiter
	splitStr := strings.Split(myString, " ")
	fmt.Println("Split String:", splitStr)

	// join a slice of strings into a single string with a specified separator
	joinedStr := strings.Join(splitStr, "-")
	fmt.Println("Joined String:", joinedStr)

	// replace the first occurrence of a substring with another substring
	replacedStrFirst := strings.Replace(myString, "Go", "Golang", 1)
	fmt.Println("Replaced String (First Occurrence):", replacedStrFirst)

	// replace all occurrences of a substring with another substring
	replacedStr := strings.ReplaceAll(myString, "Go", "Golang")
	fmt.Println("Replaced String:", replacedStr)

	// trim leading and trailing whitespace from a string
	trimmedStr := strings.TrimSpace("   Hello, Go!   ")
	fmt.Println("Trimmed String:", trimmedStr)

	// ltrim leading whitespace from a string
	ltrimmedStr := strings.TrimLeft("   Hello, Go!", " ")
	fmt.Println("Left Trimmed String:", ltrimmedStr)

	// length of a string after trimming whitespace
	trimmedLength := len(strings.TrimSpace("   Hello, Go!   "))
	fmt.Println("Length of Trimmed String:", trimmedLength)

}
