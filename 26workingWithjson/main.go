package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string   `json:"fname"`
	LastName  string   `json:"lname"`
	Hobbies   []string `json:"hobbies"`
}

func main() {

	joe := Person{FirstName: "joe", LastName: "Smith"}
	joe.Hobbies = []string{"cricket", "movies", "badminton"}

	jsonOutput, _ := json.MarshalIndent(&joe, "", "   ")

	fmt.Println(string(jsonOutput))

	joeFromJson := Person{}

	json.Unmarshal(jsonOutput, &joeFromJson)
	fmt.Println(joeFromJson)
}
