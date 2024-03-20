package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files in Golang")
	content := "This needs to go in a file"

	file, err := os.Create("./myfile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is:", length)
	defer file.Close()
	ReadFile("./myfile.txt")
}

// func readFile(fname string) {
// 	databyte, err := ioutil.ReadFile(fname)
// 	checkNilErr(err)

// 	fmt.Println("Text data inside the file is \n", databyte)
// 	fmt.Println("Text data inside the file is \n", string(databyte))
// }

func ReadFile(filename string) ([]byte, error) {
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", databyte)
	fmt.Println("Text data inside the file is \n", string(databyte))
	return os.ReadFile(filename)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
