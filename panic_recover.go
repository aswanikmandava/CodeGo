package main

import "fmt"

// recovering from panic can be done from deferred functions only

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from: ", r)
		}
	}()

	panic("Something went wrong")
	fmt.Println("This will not run")
}