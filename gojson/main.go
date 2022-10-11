package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string
	Price    int
	Platform string
	Password string
	Tags     []string
}

func main() {
	fmt.Println("Welcome to Json file in Golang")
	EncodeJson()

}

func EncodeJson() {
	mycourses := []course{
		{"Reactjs", 100, "nextgentips", "abc123", []string{"Js", "web"}},
		{"MongoDB", 500, "nextgentips", "acd123", []string{"DB", "web"}},
		{"Rabbitmq", 150, "nextgentips", "bcd123", nil},
	}

	//package data into json

	jsonfinal, err := json.MarshalIndent(mycourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", jsonfinal)

}
