package main

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	println("sum:", sum)

	names := []string{"Go", "Java", "Kotlin", "JavaScript"}
	for index, name := range names {
		println(index, name)
	}
}
