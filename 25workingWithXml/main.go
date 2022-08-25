package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	FirstName string   `xml:"fname"`
	LastName  string   `xml:"lname"`
	Hobbies   []string `xml:"hobbies"`
}

func main() {

	joe := Person{FirstName: "joe", LastName: "Smith"}
	joe.Hobbies = []string{"cricket", "movies", "badminton"}

	xmlOutput, _ := xml.MarshalIndent(&joe, "", "   ")

	fmt.Println(xml.Header)
	fmt.Println(string(xmlOutput))

	joeFromXml := Person{}

	xml.Unmarshal(xmlOutput, &joeFromXml)
	fmt.Println(joeFromXml)
}
