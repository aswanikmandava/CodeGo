package main

import (
	"fmt"
)

// function that takes a map and updates its value for a given key
// If the key does not exist, it initializes it with the given value
// map is passed by reference, so changes made inside the function will affect the original map
func updateMap(m map[string]int, key string, value int) {
	m[key] = m[key] + value
}

func main() {
	players := map[string]int{
		"Roger": 10,
		"Djcovicz":   24,
	}
	updateMap(players, "Roger", 5)
	updateMap(players, "Alice", 15) // This will initialize Alice with a value of 15
	for name, score := range players {
		fmt.Printf("%s: %d\n", name, score)
	}
}
