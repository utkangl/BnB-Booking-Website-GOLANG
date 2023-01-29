package main

import (
	"net/http"
)

const portNumber = ":8080"

// this function is a handler function

func main() {

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/about", AboutPage)

	_ = http.ListenAndServe(portNumber, nil)
}
