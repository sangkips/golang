package main

import "fmt"

func main() {
	fmt.Println("This is if else section in go")

	loginCount := 30
	var result string

	if loginCount < 5 {

		result = "This is a regular user"
	} else if loginCount <= 10 {
		result = "That is an intruder"
	} else {
		result = "Panic the system and engage safe mode"
	}
	fmt.Println(result)

	if num := 20; num < 15 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is NOT less than 10")
	}
}
