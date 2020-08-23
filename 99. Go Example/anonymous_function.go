package main

import "fmt"

func main() {
	sum := func(nums ...int) int {
		s := 0
		for _, i := range nums {
			s += i
		}
		return s
	}

	sumResult := sum(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("Sum:", sumResult)

	add := func(i int, j int) int {
		return i + j
	}

	addResult := calculate(add, 5, 5)
	fmt.Println("Add:", addResult)
}

func calculate(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}
