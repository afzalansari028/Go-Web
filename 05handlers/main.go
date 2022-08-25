package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct {
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a custom hanlder called MyHandler.")
}

func MyHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a customer handler function is called MyhandlerFunc")
}

func main() {

	http.Handle("/handler", &MyHandler{})
	http.HandleFunc("/handlerfunc", MyHandlerFunc)

	http.ListenAndServe(":8000", nil)
}
