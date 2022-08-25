package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NameHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Name: %s %s", p.ByName("firstname"), p.ByName("lastname"))
}
func main() {

	router := httprouter.New()

	router.GET("/name/:firstname/:lastname", NameHandler)

	http.ListenAndServe(":8081", router)

}
