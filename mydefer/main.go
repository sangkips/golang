package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer in golang")

	defer fmt.Println("this is defer")
	fmt.Println("This is good ")
}
