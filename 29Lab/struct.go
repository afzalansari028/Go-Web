package main

import "fmt"

type Animal struct {
	Name string
	Age  int
}
type Ocean struct {
	Name  string
	depth float32
}

type Human interface {
	Say() string
}

func (a Animal) Say() string {
	return "Hello"
}
func (o Ocean) Say() string {
	return "Water"
}

func main() {
	var h Human
	var a Animal
	var o Ocean
	// a.Name = "king"
	// a.Age = 23
	h = a

	// var a Animal
	fmt.Println(h.Say())
	h = o
	fmt.Println(h.Say())

}
