package main

import "fmt"

func main() {
	// Map 생성 및 초기화
	wordCounts := make(map[string]int)

	wordCounts["go"] = 2
	wordCounts["java"] = 15
	wordCounts["kotln"] = 5

	fmt.Println(wordCounts)
	fmt.Println("Length:",len(wordCounts))

	// Map에 Key 값이 있는지 체크
	_, isExisting := wordCounts["javascript"]
	fmt.Println(isExisting)

	// Map 요소 출력
	for key, val := range wordCounts {
		fmt.Println(key, val)
	}
}
