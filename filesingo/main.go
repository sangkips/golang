package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to file in golang")

	content := "I love the state of Nextgentips.com"

	file, err := os.Create("./logfile.txt")

	//if err != nil {
	//	panic(err)
	//}

	checknilError(err)

	length, err := io.WriteString(file, content)

	checknilError(err)

	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./logfile.txt")

}

//Create a function to read a file

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checknilError(err)

	fmt.Println("Text inside the file is \n ", string(databyte))
}

func checknilError(err error) {
	if err != nil {
		panic(err)
	}
}
