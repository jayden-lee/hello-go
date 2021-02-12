package main

import "fmt"

func main() {
	// 변수명 [배열 크기] 데이터 타입
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Println("Length:", len(numbers))

	result := 0
	for _, value := range numbers {
		result += value
	}
	fmt.Println("Sum:", result)

	// 다차원배열
	var multiArray [2][3]int
	multiArray[0][0] = 0
	multiArray[0][1] = 1
	multiArray[0][2] = 2
	multiArray[1][0] = 3
	multiArray[1][1] = 4
	multiArray[1][2] = 5
	fmt.Println(multiArray)

}
