package main

import (
	"fmt"
)

func main() {
	// declaring a map with string keys and int values
	ages := make(map[string]int)
	// adding elements to the map
	ages["Alice"] = 30
	ages["Bob"] = 25
	ages["Charlie"] = 35
	// accessing elements from the map
	fmt.Println("Alice's age:", ages["Alice"])
	fmt.Println("Bob's age:", ages["Bob"])
	fmt.Println("Charlie's age:", ages["Charlie"])

	// declaring and initializing a map with values
	scores := map[string]int{
		"Math":    90,
		"Science": 85,
		"English": 95,
	}
	fmt.Println("Scores:", scores)

	// iterating over a map using range
	for subject, score := range scores {
		fmt.Printf("Subject: %s, Score: %d\n", subject, score)
	}
	fmt.Printf("number of items in map: ", len(scores))

	// overwriting the value of a key in the map
	scores["Math"] = 95
	fmt.Println("Updated Scores of Math:", scores["Math"])

	// deleting an element from the map
	delete(scores, "Science")
	fmt.Println("Scores after deleting Science:", scores)
}