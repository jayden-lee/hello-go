package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	r := Rectangle{10., 20.}
	c := Circle{10}

	showArea(r, c)
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		area := s.area()
		fmt.Println("area:", area)
	}
}
