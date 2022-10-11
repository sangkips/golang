package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["R"] = "Ruby"
	languages["P"] = "Python"
	languages["G"] = "Golang"
	languages["J"] = "Java"

	fmt.Println("List of languages:", languages)

	delete(languages, "J")
	fmt.Println(languages)

	//Looping in maps

	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)
	}
}
