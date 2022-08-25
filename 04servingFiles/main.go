package main

import (
	"log"
	"net/http"
)

func main() {
	// fs := http.FileServer(http.Dir("/Users/1026795mgr/go/src/Go-Web/04servingFiles"))

	// http.Handle("/projectfiles/", http.StripPrefix("/projectfiles", fs))

	http.HandleFunc("/samplepdf", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "pdf/sample.pdf")
	})

	log.Print("Server started...!")

	http.ListenAndServe(":9000", nil)
}
