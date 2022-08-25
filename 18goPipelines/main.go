package main

import (
	"net/http"
	"strings"
	"text/template"
)

type Person struct {
	FirstName string
	LastName  string
}

func removeSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

func main() {

	templateFile := "C:/Users/1026795mgr/go/src/Go-Web/18goPipelines/personlist.html"

	fmap := template.FuncMap{
		"removeSpaces": removeSpaces,
		"capitalize":   strings.ToUpper,
	}

	http.HandleFunc("/shortlist/", func(w http.ResponseWriter, r *http.Request) {
		tp, _ := template.New("personlist.tmpl").Funcs(fmap).ParseFiles(templateFile)

		people := []Person{
			{"Bo  b", "Barke   r"},
			{"Lar  ry", "F   lint"}}

		tp.Execute(w, people)
	})

	http.HandleFunc("/longlist/", func(w http.ResponseWriter, r *http.Request) {
		tp, _ := template.New("personlist.tmpl").Funcs(fmap).ParseFiles(templateFile)

		people := []Person{
			{"Bo  b", "Barke   r"},
			{"Lar  ry", "F   lint"},
			{"Akb  ar", "Hu  ssain"},
			{"Laal  a", "Raj   put"},
			{"Rajkum   a r", "Ra o"}}

		tp.Execute(w, people)
	})

	http.ListenAndServe(":8082", nil)
}
