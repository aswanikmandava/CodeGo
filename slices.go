package main

import (
	"fmt"
	"slices"
)

func main() {

	// slices are similar to arrays but they are more flexible and can grow and shrink in size. Slices are built on top of arrays and provide a more convenient way to work with sequences of data.

	// declaring a slice of integers
	var numbers []int
	// in arrays, we have to specify the size of the array, but in slices, we don't have to specify the size.
	// appending values to the slice
	numbers = append(numbers, 10)
	numbers = append(numbers, 20)
	numbers = append(numbers, 30)
	fmt.Println("numbers: ", numbers)

	numbers2 := []int{1, 2, 3, 4, 5}
	fmt.Println("numbers2: ", numbers2)
	numbers2 = append(numbers2, 6, 7, 8)
	fmt.Println("numbers2 after appending: ", numbers2)

	// removing an element from a slice using element index
	var elementToRemove int = numbers2[3] // element at index 3 is 4
	numbers2 = append(numbers2[:3], numbers2[4:]...)
	fmt.Println("numbers2 after removing element", elementToRemove, "is", numbers2)

	// another way to remove an element from a slice
	numbers2 = slices.Delete(numbers2, 3, 4) // remove elements from index 3 to 4

	// using Delete function from slices package to remove an element from a slice
	numbers2 = slices.DeleteFunc(numbers2, func(x int) bool {
		return x == 2 // remove element
	})

	fmt.Println("numbers2 after removing element 2 is", numbers2)

}
