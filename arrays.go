package main

import (
	"fmt"
)

func main() {
	// declaring an array of integers with size 5
	var arr [5]int
	// assigning values to the array
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	fmt.Println("arr: ", arr)
	// declaring and initializing an array of strings
	names := [3]string{"Alice", "Bob", "Charlie"}
	fmt.Println("names: ", names)
	// declaring and initializing an array of floats
	prices := [4]float64{10.5, 20.0, 30.75, 40.25}
	fmt.Println("prices: ", prices)
	// declaring and initializing an array of booleans
	flags := [3]bool{true, false, true}
	fmt.Println("flags: ", flags)
	// declaring and initializing an array of structs
	type Person struct {
		name string
		age  int
	}
	people := [2]Person{{name: "John", age: 30}, {name: "Alice", age: 25}}
	fmt.Println("people: ", people)
	// get first person from people array
	p1 := people[0]
	p2 := people[1]
	fmt.Println("Comparing p1 and p2: %t", p1==p2)
}
