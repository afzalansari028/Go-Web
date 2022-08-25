package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/notfound/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	})

	http.HandleFunc("/servererror/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})

	http.HandleFunc("/customheader/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Custom-Header", "This is my custom header.")
		io.WriteString(w, "Check out the custom header using F12.")
	})
	log.Print("server started....!")
	http.ListenAndServe(":8080", nil)
}
