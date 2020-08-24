package main

import "fmt"

func main() {
	channel := make(chan int)

	go func() {
		channel <- 1234
	}()

	var i int
	i = <- channel
	fmt.Println(i)
}
