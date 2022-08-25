package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	// response, _ := http.Get("https://postman-echo.com/get")
	response, _ := http.Get("https://go.dev/")

	responseBody, _ := ioutil.ReadAll(response.Body)

	fmt.Println("Code: ", response.StatusCode)
	fmt.Println(string(responseBody))

}
