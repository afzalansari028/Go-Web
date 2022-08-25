package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/afzalansari028/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// routes.RegisterBookStoreRoutes(r)
	r.HandleFunc("/book", controllers.GetBook).Methods("GET")
	fmt.Println("main function")
	// http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8082", r))
}

// func GetBook(w http.ResponseWriter, r *http.Request) {
// 	log.Print("Get book controller")
// 	newBooks := models.GetAllBooks()
// 	res, _ := json.Marshal(newBooks)
// 	w.Header().Set("Content-Type", "json/application")
// 	w.WriteHeader(http.StatusOK)
// 	// w.Write(res)
// }
