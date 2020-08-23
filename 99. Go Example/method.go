package main

import "fmt"

type Rect struct {
	width  int
	height int
}

func (r Rect) area() int {
	return r.width * r.height
}

func main() {
	rect := Rect{10, 20}
	area := rect.area()
	fmt.Println(area)
}
