package main

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	FirstName string
	LastName  string
}

type PersonJsonFields struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
}

type PersonJsonMarshaler struct {
	FirstName string
	LastName  string
}

func (p *PersonJsonMarshaler) MarshalJSON() ([]byte, error) {
	return []byte("{\"Name\":\"" + p.FirstName + " " + p.LastName + "\"}"), nil
}

func main() {

	http.HandleFunc("/json/default/", func(w http.ResponseWriter, r *http.Request) {
		j, _ := json.Marshal(&Person{FirstName: "Atif", LastName: "Aslam"})
		w.Write(j)
		// r.Header.Set("Content-Type")
	})

	http.HandleFunc("/json/fields/", func(w http.ResponseWriter, r *http.Request) {
		j, _ := json.Marshal(&PersonJsonFields{FirstName: "Atif", LastName: "Aslam"})
		w.Write(j)
	})

	http.HandleFunc("/json/marshaler/", func(w http.ResponseWriter, r *http.Request) {
		j, _ := json.Marshal(&PersonJsonMarshaler{FirstName: "Atif", LastName: "Aslam"})
		w.Write(j)
	})

	http.ListenAndServe(":8080", nil)
}
