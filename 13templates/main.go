package main

import (
	"log"
	"net/http"
	"text/template"

	httprouter "github.com/julienschmidt/httprouter"
)

type Person struct {
	FirstName string
	LastName  string
}

func main() {

	temp1 := "<B><font color='green'>First Name : </font></B> {{.FirstName}} <br> <B><font color='red'>Last Name : </font></B> {{.LastName}}"
	// temp2 := "<B><font color='Blue'>Age : </font></B> {{.Age}}"

	router := httprouter.New()

	router.GET("/name/:firstName/:lastName", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		t := template.New("NameTemplate")
		tp, _ := t.Parse(temp1)
		n := t.Name()
		log.Print(n)
		tp.Execute(w, &Person{ps.ByName("firstName"), ps.ByName("lastName")})
	})
	http.ListenAndServe(":8081", router)
}
