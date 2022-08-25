package main

import (
	"fmt"
	"log"
	"net/http"
)

type PersonHandler struct {
	FirstName string
	LastName  string
}

type CarHandler struct {
	Make  string
	Model string
	Year  int
}

func (h *PersonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Name: %v %v", h.FirstName, h.LastName)
}

func (h *CarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Make: %v -- Model: %v -- Year: %v", h.Make, h.Model, h.Year)
}

func defaultHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "This is default handler...!")
}

func main() {

	http.Handle("/person", &PersonHandler{"Afzal", "Ansari"})
	http.Handle("/car", &CarHandler{"Ford", "Fiesta", 2005})

	http.HandleFunc("/", defaultHandler)

	log.Print("server started at port : 8282")
	http.ListenAndServe(":8282", nil)
}
