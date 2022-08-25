package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/unrecognized_path.html")
	})

	http.HandleFunc("/demopath1/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/demopath1.html")
	})

	http.HandleFunc("/demopath1/subpath", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/demopath1_subpath.html")
	})

	http.HandleFunc("/demopath2/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/demopath2.html")
	})

	log.Print("Server started...")

	http.ListenAndServe(":8082", nil)

}
