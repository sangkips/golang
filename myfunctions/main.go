package main

import "fmt"

func main() {
	fmt.Println("Welcome to golang functions")

	total := plusPlus(10, 20)

	fmt.Println("The result is: ", total)

	result := adder(50, 20, 40, 85)
	fmt.Println("The results are: ", result)
}

func plusPlus(valone, valtwo int) int {

	return valone + valtwo

}

func adder(values ...int) int {
	mytotal := 0

	for _, val := range values {
		mytotal += val
	}

	return mytotal

}
