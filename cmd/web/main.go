package main

import (
	"fmt"
	"net/http"

	"github.com/utkangl/GoWEB/pkg/funcs"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", funcs.HomePage)
	http.HandleFunc("/about", funcs.AboutPage)
	fmt.Println("Starting the Application on Port: ", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
