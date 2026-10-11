package main

import (
	"fmt"
)

type Shape2 interface {
	getArea() float64
}

type Square struct {
	sideLength float64
}

type Rectangl struct {
	base float64
	height float64
}

func main() {
	s := Square{sideLength: 10.0}
	r := Rectangl{base: 1.5, height: 20}
	printArea(s)
	printArea(r)

}

func (sq Square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (r Rectangl) getArea() float64 {
	return 0.5 * r.base * r.height
}

func printArea(s Shape2) {
	fmt.Println(s.getArea())
}