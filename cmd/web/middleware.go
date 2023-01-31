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

// adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {

	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
