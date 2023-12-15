package main

import "fmt"

const LokinToken string = "fdgdg" // public

func main() {
	var username string = "Harsh"
	fmt.Println("Variables")
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n", smallValue)

	var smallFloat float32 = 255.324234789724980
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	var msg = "Hello"
	fmt.Println(msg)

	// no var style but not allowed outside func.
	noOfUser := 3000
	fmt.Println(noOfUser)

	fmt.Println(LokinToken)
}
