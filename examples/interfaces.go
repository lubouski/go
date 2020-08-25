package main

import (
	"math"
	"fmt"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width  float64
	height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	c1 := circle{4.5}
	r1 := rect{5, 8}
	shapes := []shape{c1, r1}

	/*	but we don't know the type of circle and rect functions

		shapes := []type{c1, r1}
		shape[0].area()
		shape[1].area()
	*/

	for _, shape := range shapes {
		fmt.Println(shape.area())
	}
}
