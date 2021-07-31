package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)

	loginPw := `password123`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPw))
	if err != nil {
		fmt.Println("YOU CANT LOGIN!", err)
	}
}
