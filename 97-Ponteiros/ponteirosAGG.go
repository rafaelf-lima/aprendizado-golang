package main

import "fmt"


type User struct {
	email string
	username string
	age int
}


func (u User) Email() string {
	return u.email
}

func (u *User) updateEmail(email string) {
	u.email = email
}

func main(){
	user := User{
		email: "sdsdaa@foo.com",
	}
	fmt.Println(user.Email())
	user.updateEmail("sdasdasdas@bar.com")
	fmt.Println(user.email)
}