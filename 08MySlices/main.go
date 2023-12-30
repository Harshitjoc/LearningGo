package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in Golang")
	var fruitList = []string{"Apple", "Mango", "Peach"}
	// fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Tomato", "Banana")
	// fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 944
	highScore[2] = 264
	highScore[3] = 874
	// highScore[4] = 777

	highScore = append(highScore, 555, 666, 321)

	// fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)
	// fmt.Println(sort.IntsAreSorted(highScore))

	// how to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
