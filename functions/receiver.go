package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// defining a receiver method for the Person struct
// The getMetadata method returns a string representation of the Person struct
func (p Person) getMetadata() string {
	return fmt.Sprintf("Person{Name: %s, Age: %d}", p.Name, p.Age)
}

type Rectangle struct {
	Width  float32
	Height float32
}

// defining a receiver method for the Rectangle struct
// The Area method returns the area of the rectangle
func (r Rectangle) Area() float32 {
	return r.Width * r.Height
}

func main() {
	person := Person{Name: "Alice", Age: 30}
	fmt.Println(person.getMetadata())
	shape := Rectangle{Width: 5, Height: 10}
	fmt.Printf("Area of rectangle: %.2f\n", shape.Area())
}
