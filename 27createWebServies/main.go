package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
}

func main() {

	var people []Person

	http.HandleFunc("/person/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet: //USING GET REQUEST METHOD

			jsonRequest, _ := json.MarshalIndent(&people, " ", "   ")
			fmt.Fprintf(w, string(jsonRequest))

		case http.MethodPost: //USING POST REQUEST METHOD

			requestBodyBytes, _ := ioutil.ReadAll(r.Body)
			var newPerson Person
			json.Unmarshal(requestBodyBytes, &newPerson)

			for i := range people {
				if people[i].ID == newPerson.ID {
					w.WriteHeader(http.StatusConflict)
					return
				}
			}
			people = append(people, newPerson)

		case http.MethodPut: //USING PUT REQUEST METHOD

			requestBodyBytes, _ := ioutil.ReadAll(r.Body)
			var newPerson Person
			json.Unmarshal(requestBodyBytes, &newPerson)

			for i := range people {
				if people[i].ID == newPerson.ID {
					people[i].FirstName = newPerson.FirstName
					people[i].LastName = newPerson.LastName
					return
				}
			}
			w.WriteHeader(http.StatusNotFound)

		case http.MethodDelete: //USING DELETE METHOD

			requestBodyBytes, _ := ioutil.ReadAll(r.Body)
			var newPerson Person
			json.Unmarshal(requestBodyBytes, &newPerson)

			for i := range people {
				if people[i].ID == newPerson.ID {
					copy(people[i:], people[i+1:])
					people = people[:len(people)-1]
					return
				}
			}
			w.WriteHeader(http.StatusNotFound)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return

		}
	})

	http.ListenAndServe(":8082", nil)

}
