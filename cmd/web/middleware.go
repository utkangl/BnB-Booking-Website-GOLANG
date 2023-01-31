package main

import (
	"fmt"
	"net/http"
)

// simple middleware that lets us know that page has been visited
func WriteToConsole(next http.Handler) http.Handler {

	return http.HandlerFunc(func(Res http.ResponseWriter, Req *http.Request) {

		fmt.Println("Page has been visited")
		next.ServeHTTP(Res, Req)
	})
}
