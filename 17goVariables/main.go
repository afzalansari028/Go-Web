package main

import (
	"html/template"
	"net/http"
)

type Person struct {
	Id        int
	FirstName string
	LastName  string
}

func main() {

	templateFile := "C:/Users/1026795mgr/go/src/Go-Web/17goVariables/personlist.tmpl"

	http.HandleFunc("/shortlist/", func(w http.ResponseWriter, r *http.Request) {

		tp, _ := template.ParseFiles(templateFile)

		people := []Person{
			{1, "Afzal", "Ansari"},
			{2, "Praveen", "Yadav"}}
		tp.Execute(w, people)
	})

	//this template for the list of people
	http.HandleFunc("/longlist/", func(w http.ResponseWriter, r *http.Request) {

		tp, _ := template.ParseFiles(templateFile)

		people := []Person{
			{10, "Sachin", "Pradhan"},
			{20, "Shubham", "Jadhav"},
			{30, "Usman", "Ali"},
			{40, "Ajeet", "Yadav"},
			{50, "Shikhar", "Dhavan"},
			{60, "Atif", "Aslam"},
			{70, "Gulzar", "Ahmed"},
			{80, "Faizan", "Ansari"}}
		tp.Execute(w, people)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tp, _ := template.ParseFiles(templateFile)

		tp.Execute(w, nil)
	})

	http.ListenAndServe(":8080", nil)
}
