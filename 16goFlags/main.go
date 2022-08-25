package main

import (
	"flag"
	"net/http"
)

func main() {

	var port, message string

	flag.StringVar(&port, "port", "8080", "The HTTP port")
	flag.StringVar(&message, "message", "", "The response message")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte(message))
	})

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}

}
