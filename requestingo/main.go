package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to golang")
	//PrintGetRequest()
	//PrintPostRequest()
	PerformPostFormRequest()
}

func PrintGetRequest() {
	const myurl = "http://localhost:1111/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Status code is: ", response.StatusCode)
	//fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	bytecount, _ := responseString.Write(content)
	fmt.Println(responseString.String())
	fmt.Println(bytecount)
}

func PrintPostRequest() {
	const myurl = "http://localhost:1111/post"

	// create json payload

	requestBody := strings.NewReader(`
		{
			"Fname":"Kipkoech",
			"Lname":"Sang",
			"Gender":"Male"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:1111/postform"

	data := url.Values{}

	data.Add("fname", "Sang")
	data.Add("lname", "Kips")
	data.Add("gender", "male")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
