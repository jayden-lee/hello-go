package main

import "fmt"

// type 키워드, 사용자 정의 타입 이름, 기본 타입
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
	changeUserName(&user2)
	fmt.Println(user2)

	user3 := new(user)
	fmt.Println(user3)
}

func changeUserName(u *user) {
	u.id = "new id"
	u.name = "new name"
}
