package main

// multiple imports
import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	// main() is one per project, and is the entry point for the program
	fmt.Println("Hello, World!")
	fmt.Println("Random number:", rand.Intn(100))
	fmt.Println("Pi: ", math.Pi)
}