package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for pizza:")

	// comma ok || comma err

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating,", input)
	fmt.Printf("Type of rating is %T", input)
}
