package main

import "fmt"

// Go interfaces look exclusively for receiver functions. 
// A standalone single-argument function cannot be used to satisfy an interface.

// any type inside of our program that implements a function called greet() that returns a string.
// And that's true of both engSite and hindiSite will automatically promoted as members of Greeter type
type Greeter interface {
	greet() string	// function signature that includes the name of the function, the return type, and the parameters
}

type engSite struct{}
type hindiSite struct{}

func (e engSite) greet() string {
	return "Hi!"
}

func (h hindiSite) greet() string {
	return "Namaste!"
}

// 2. Fix the duplication! Accept the Greeter interface instead of a concrete type.
// This function can now accept engSite, hindiSite, or ANY type that has a greet() method.
func printGreeting(g Greeter) {
	fmt.Println(g.greet())
}

func main() {
	e := engSite{}
	h := hindiSite{}

	// The exact same function now handles both types perfectly
	printGreeting(e)
	printGreeting(h)
}
