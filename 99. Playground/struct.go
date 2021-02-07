package main

import "fmt"

type user struct {
	id   string
	name string
}

func main() {
	user1 := user{}
	user1.id = "jayden-lee"
	user1.name = "jayden"
	fmt.Println(user1)

	user2 := user{"id", "name"}
	fmt.Println(user2)

	user3 := new(user)
	fmt.Println(user3)
}
