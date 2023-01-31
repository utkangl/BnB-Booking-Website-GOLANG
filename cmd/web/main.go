package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/utkangl/GoWEB/packages/config"
	"github.com/utkangl/GoWEB/packages/handlers"
	"github.com/utkangl/GoWEB/packages/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // uses https insted of http

	app.Session = session

	tempCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tempCache

	repository := handlers.CreateRepo(&app)
	handlers.SetRepo(repository)

	render.SetConfig(&app)

	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	fmt.Println("Starting the Application on Port: ", portNumber)

	err = serve.ListenAndServe()
	log.Fatal(err, "Cannot serve")

}
