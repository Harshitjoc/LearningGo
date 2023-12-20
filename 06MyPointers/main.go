package main

import "fmt"

func main() {
	fmt.Println("Learning About Pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is", ptr)

	myNumber := 23

	var ptr = &myNumber // reference means &

	fmt.Println("Value of memory address is", ptr)
	fmt.Println("Value of actual pointer is", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New Value is :", myNumber)
}
