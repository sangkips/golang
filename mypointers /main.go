package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers in Golang")

	//var ptr *int
	//fmt.Println(ptr)

	total := 50

	var ptr = &total
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = *ptr / 2
	fmt.Println(*ptr)
}
