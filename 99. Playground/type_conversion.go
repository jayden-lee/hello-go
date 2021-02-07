package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number int = 3
	fmt.Println(reflect.TypeOf(number))
	fmt.Println(reflect.TypeOf(float64(number)))
}
