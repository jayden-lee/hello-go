package main

func main() {
	// const 상수명 상수 타입
	const number int = 1000000
	println("number:", number)

	// 상수 타입 추론
	const prefix = "GO"
	println("prefix:", prefix)

	const (
		Apple  = iota // 0
		Banana        // 1
		Grape         // 2
		Orange        // 3
	)

	println("Apple", Apple)
	println("Banana", Banana)
	println("Grape", Grape)
	println("Orange", Orange)
}
