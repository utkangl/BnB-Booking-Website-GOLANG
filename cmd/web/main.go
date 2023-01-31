package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/utkangl/GoWEB/packages/config"
	"github.com/utkangl/GoWEB/packages/handlers"
	"github.com/utkangl/GoWEB/packages/render"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig

	tempCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tempCache

	repository := handlers.CreateRepo(&app)
	handlers.SetRepo(repository)

	render.SetConfig(&app)

	// http.HandleFunc("/", handlers.Repo.HomePage)
	// http.HandleFunc("/about", handlers.Repo.AboutPage)
	// _ = http.ListenAndServe(portNumber, nil)

	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	fmt.Println("Starting the Application on Port: ", portNumber)

	err = serve.ListenAndServe()
	log.Fatal(err, "Cannot serve")

}
