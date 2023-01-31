package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// simple middleware that lets us know that page has been visited
func WriteToConsole(next http.Handler) http.Handler {

	return http.HandlerFunc(func(Res http.ResponseWriter, Req *http.Request) {

		fmt.Println("Page has been visited")
		next.ServeHTTP(Res, Req)
	})

}

// little middleware to protect from csrf attacks
func NoSurf(next http.Handler) http.Handler {

	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}
