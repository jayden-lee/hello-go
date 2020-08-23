package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(f.Name())
}
