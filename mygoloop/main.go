package main

import "fmt"

func main() {
	fmt.Println("Looping in golang")

	//days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Sunday"}
	//fmt.Println(days)
	// This is really used
	//for i := 0; i < len(days); i++ {
	//	fmt.Println(days[i])
	//}

	//this is mostly used

	//for i := range days {
	//	fmt.Println(days[i])
	//}

	//for _, day := range days {
	//fmt.Printf("Today day is %v\n", day)
	//}

	myValue := 0

	for myValue < 10 {

		if myValue == 3 {
			goto val
		}

		if myValue == 5 {
			myValue++
			continue
		}
		fmt.Println("Value is ", myValue)
		myValue++
	}

val:
	fmt.Println("Run your race")
}
