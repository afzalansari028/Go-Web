package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		// fmt.Printf("r: %v\n", r)
	})

	http.HandleFunc("/withargs/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is a response from a %s request. The arguments are: arg1:%s, arg2:%s ", r.Method, r.URL.Query().Get("arg1"), r.URL.Query().Get("arg2"))
		r.URL.Query().Get("arg1")
	})

	http.HandleFunc("/postonlywithbody/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		body, _ := ioutil.ReadAll(r.Body)

		fmt.Fprintln(w, string(body))
	})

	http.ListenAndServe(":8080", nil)
}
