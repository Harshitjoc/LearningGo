package main

import "fmt"

func main() {

	// there is no class in golang

	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	harsh := User{"Harsh", "harsh@go.dev", true, 16}
	fmt.Println(harsh)
	fmt.Printf("Harsh details are: %+v\n", harsh)
	fmt.Printf("Name is %v and Email is %v\n", harsh.Name, harsh.Email)
	harsh.GetStatus()
	harsh.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User Active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}
