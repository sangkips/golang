package main

import "fmt"

func main() {
	fmt.Println("This is variables in Golang")

	var myname string = "Kipkoech Sang" //variable declaration using var
	fmt.Println(myname)

	const smalldigit int = 20 // variable declaration using const.
	fmt.Println(smalldigit * 2)
	fmt.Println(smalldigit + 50)
	fmt.Println(smalldigit / 2)
	fmt.Println(smalldigit % 2)

	var myfloat float32 = 50.5
	fmt.Println(myfloat)
	fmt.Println(myfloat * 2)
	fmt.Println(myfloat + 50)
	fmt.Println(myfloat / 2)

	myvar := 2 //without var declaration
	fmt.Println(myvar)

}
