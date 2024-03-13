package main

import "fmt"

func main() {
	fmt.Println("Functions in Golang")
	greeter()
	greeterTwo()

	result := adder(65, 45)
	fmt.Println("Result is:", result)

	proRes, myMessage := proAdder(2, 5, 3, 6)
	fmt.Println("Pro Result is:", proRes)
	fmt.Println("Pro Message is:", myMessage)
}

func greeter() {
	fmt.Println("Namstey from golang")
}
func greeterTwo() {
	fmt.Println("Another method")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi Pro result function"
}
