package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/utkangl/GoWEB/pkg/config"
	"github.com/utkangl/GoWEB/pkg/funcs"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig

	tempCache, err := funcs.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tempCache

	http.HandleFunc("/", funcs.HomePage)
	http.HandleFunc("/about", funcs.AboutPage)

	fmt.Println("Starting the Application on Port: ", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
