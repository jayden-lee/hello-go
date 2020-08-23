package main

import "fmt"

func main() {
	msg := "Hello go"
	printlnMessage(msg)

	changeMessage(&msg)
	printlnMessage(msg)

	printMultipleMessage("Java", "Kotlin", "Go")

	result := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Sum:", result)

	count, total := sumAndCount(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("Count:", count, "Sum:", total)
}

// 반환형 없는 함수
func printlnMessage(msg string) {
	fmt.Println(msg)
}

// 주소 값을 받아서 문자열 값 수정
func changeMessage(msg *string) {
	*msg = "Change Message"
}

// 가변 인자 함수
func printMultipleMessage(msg ...string) {
	for _, s := range msg {
		fmt.Print(s + " ")
	}
	fmt.Println()
}

// 반환형 있는 함수
func sum(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

// 복수 개의 값을 반환 하는 함수
func sumAndCount(nums ...int) (int, int) {
	s := 0     // 합계
	count := 0 // 개수
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}
