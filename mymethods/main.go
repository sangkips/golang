package main

import "fmt"

func main() {
	myname := User{"Sang", 25, true, "kips@gmail.com"}
	fmt.Println(myname)
	fmt.Printf("My details are: %+v\n", myname)
	fmt.Printf("My name is: %v and my email is %v\n", myname.Name, myname.Email)

	myname.GetStatus()
	myname.NewEmail()
	fmt.Printf("My name is: %v and my email is %v\n", myname.Name, myname.Email)
}

type User struct {
	Name   string
	Age    int
	Status bool
	Email  string
}

func (u User) GetStatus() {
	fmt.Println("Does the user exist? ", u.Status)
}

func (u User) NewEmail() {
	u.Email = "sang@gmail.com"

	fmt.Println(u.Email)
}
