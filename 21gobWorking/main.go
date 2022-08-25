package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Person struct {
	FirstName string
	Lastname  string
}

func writeBinaryFile(data interface{}, file string) {
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)
	encoder.Encode(data)
	ioutil.WriteFile(file, buf.Bytes(), 0600)
}

func readBinaryFile(data interface{}, file string) {
	raw, _ := ioutil.ReadFile(file)
	buf := bytes.NewBuffer(raw)
	decoder := gob.NewDecoder(buf)
	decoder.Decode(data)
}

func main() {
	file := "C:/Users/1026795mgr/go/src/Go-Web/21gobWorking/person.txt"

	person := Person{"Bob", "Barker"}

	writeBinaryFile(person, file)

	var readPerson Person
	readBinaryFile(&readPerson, file)

	fmt.Println(readPerson)

}
