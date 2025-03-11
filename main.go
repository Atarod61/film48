package main

import (
	"film48/user"
	"fmt"
)

func main() {
	user1 := user.User{ID: 1, Name: "alireza", Email: "alireza@gmail.com"}
	fmt.Println(user1)
	user2 := user.SerUser2(1, "alireza", "alireza@gmail.com")
	fmt.Println(user2)
}
