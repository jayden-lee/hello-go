package main

import (
	"fmt"
)

func main() {
	// 슬라이스 생성
	s := make([]int, 5, 10)
	fmt.Println(len(s), cap(s))

	// 슬라이스 출력
	s2 := []string {"a", "b", "c"}
	for _, value := range s2 {
		fmt.Println("value:", value)
	}

	// 부분 슬라이스
	scores := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	subSlice := scores[3:6]
	fmt.Println(subSlice)

	// 슬라이스 추가
	scores = append(scores, 10)
	fmt.Println(scores)
}
