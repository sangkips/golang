package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://nileotech.com"

func main() {
	fmt.Println("Go web request")

	response, err := http.Get(url)

	checknilError(err)

	fmt.Printf("Response type is: %T\n", response)

	defer response.Body.Close()

	databyte, err := ioutil.ReadAll(response.Body)

	checknilError(err)

	content := string(databyte)

	fmt.Println(content)
}

func checknilError(err error) {
	if err != nil {
		panic(err)
	}
}
