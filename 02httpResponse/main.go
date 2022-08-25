package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

// func main() {
// 	testHTTPCall := func(responseWriter http.ResponseWriter, request *http.Request) {
// 		io.WriteString(responseWriter, "This is a test response")
// 	}

// 	request := httptest.NewRequest("GET", "http://testurl.com", nil)
// 	recorder := httptest.NewRecorder()
// 	testHTTPCall(recorder, request)

// 	response := recorder.Result()
// 	responseBody, _ := io.ReadAll(response.Body)

// 	fmt.Println(response.StatusCode)
// 	fmt.Println(string(responseBody))
// }

func main() {
	testServer := httptest.NewServer(http.HandlerFunc(func(responseWriter http.ResponseWriter, r *http.Request) {
		responseWriter.WriteHeader(http.StatusNotFound)
		io.WriteString(responseWriter, "This is a test response.")
	}))
	defer testServer.Close()

	response, _ := http.Get(testServer.URL)

	reseponseBody, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(string(reseponseBody))
}
