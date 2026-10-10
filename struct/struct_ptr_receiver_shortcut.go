package main

import "fmt"

type Shape struct {
	name string
	area float32
}

// pointer receiver function
func (sPtr *Shape) setArea(v float32) {
	(*sPtr).area = v
}

func (s Shape) printDetails() {
	fmt.Printf("%+v", s)
}

func main() {

	circle := Shape{
		name: "Circle",
		area: 0.10,
	}
	fmt.Println("------- Initial details ------")
	circle.printDetails()

	// Go will automatically pass pointer to setArea function even though circle is of type Shape but, not a pointer to type Shape
	// This is a shortcut that Go allows
	// we don't have to define as following
	// 	cPtr := &circle
	// 	cPtr.setArea(100.05)
	circle.setArea(100.05)
	fmt.Println("\nUpdated area. Printing details ....")
	circle.printDetails()

}
