package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Mango", "Pears", "Quava", "Apple"}
	fmt.Println(fruitList)

	fruitList = append(fruitList, "Banana", "Tomato")

	fmt.Println(fruitList)

	// This is very important concept in slices, its used alot

	fruitList = append(fruitList[:5])
	fmt.Println(fruitList)

	highscores := make([]int, 5)

	highscores[0] = 124
	highscores[1] = 500
	highscores[2] = 254
	highscores[3] = 365
	highscores[4] = 700

	highscores = append(highscores, 1000, 5000, 750)
	fmt.Println(highscores)

	sort.Ints(highscores)
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	//Removing a value from slice based on index

	var courses = []string{"React", "Javascript", "Rust", "Golang", "Python", "Ruby"}

	fmt.Println(courses)

	var index int = 1

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

	sea := []string{"victoria", "india"}
	fmt.Println(sea)

	sea = append(sea, "Nakuru")
	fmt.Println(sea)

}
