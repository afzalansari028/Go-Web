package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

type Group struct {
	Name   string
	People []Person
}

func main() {

	person1 := Person{"George", "Smith"}
	person2 := Person{"Afzal", "Ansari"}
	person3 := Person{"Lucky", "Brown"}
	person4 := Person{"Ronald", "Duthers"}

	people1 := []Person{person1, person2}
	people2 := []Person{person3, person4}

	gp := []Group{
		{Name: "Accounting", People: people1},
		{Name: "Sales", People: people2}}

	fmt.Println(gp)                       //display the entire array of groups
	fmt.Println(gp[0])                    //display the first group in the array
	fmt.Println(gp[1])                    //display the second group in the array
	fmt.Println(gp[1].People[0])          //display the first person in the second (Sales) group
	fmt.Println(gp[1].People[0].LastName) //display the lastname of the first person in the second group
}
