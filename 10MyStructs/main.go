package main

import "fmt"

func main() {

	// there is no class in golang

	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	harsh := User{"Harsh", "harsh@go.dev", true, 16}
	fmt.Println(harsh)
	fmt.Printf("Harsh details are: %+v\n", harsh)
	fmt.Printf("Name is %v and Email is %v", harsh.Name, harsh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
