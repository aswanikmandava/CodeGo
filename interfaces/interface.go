package main

import (
	"fmt"
)

// defining an interface named Shape
// The Shape interface has a method Area that returns a float32 value
type Shape interface {
	Area() float32
}

type Rectangle struct {
	Width  float32
	Height float32
}

type Circle struct {
	Radius float32
}

// defining a receiver method for the Rectangle struct
func (r Rectangle) Area() float32 {
	return r.Width * r.Height
}

// defining a receiver method for the Circle struct
func (c Circle) Area() float32 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	// creating instances of Rectangle and Circle
	rectangle := Rectangle{Width: 5, Height: 10}
	circle := Circle{Radius: 7}

	// using the Shape interface
	shapes := []Shape{rectangle, circle}
	for _, shape := range shapes {
		fmt.Printf("Area: %.2f\n", shape.Area())
	}
}
