package main

import "fmt"

func main() {
	var sum int // variable declaration with default value 0
	for i := 0; i < 5; i++ {
		sum += i
		fmt.Println("Current value of i:", i, "Current sum:", sum)
	}
	fmt.Println("Result: ", sum)

	persons := []string{"Alice", "Bob", "Charlie", "David", "Eve"}
	for index, name := range persons {
		fmt.Printf("Index: %d, Name: %s\n", index, name)
	}

	var count int = 0

	for {
		if count == 4 {
			break
		}
		fmt.Println("Count=", count)
		count++
	}

	for i:=1; i<5; i++ {
		if i == 2 {
			continue // skip the rest of the block
		}
		fmt.Println("i=", i)
	}
}
