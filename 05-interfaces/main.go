package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
}

type Square struct {
	length, height float64
}

func (s Square) area() float64 {
	return s.length * s.height
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	s := Square{10, 20}
	printArea(s)

	c := Circle{40}
	printArea(c)
}

func printArea(g Geometry) {
	fmt.Printf("This figure has an area of %v \n", g.area())
}
