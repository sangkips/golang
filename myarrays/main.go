package main

import (
	"fmt"
)

func main() {
	var fruitList [4]string

	fruitList[0] = "Orange"
	fruitList[1] = "Onion"
	fruitList[2] = "Tomatoes"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("total fruit is: ", len(fruitList))

	var vegList = [3]string{"skuma", "Managu", "kunde"}
	fmt.Println("veg list is: ", vegList)
	fmt.Println("veg list: ", len(vegList))

}
