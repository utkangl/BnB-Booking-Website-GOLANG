package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/utkangl/GoWEB/internalPackages/config"
	"github.com/utkangl/GoWEB/internalPackages/handlers"
	"github.com/utkangl/GoWEB/internalPackages/models"
	"github.com/utkangl/GoWEB/internalPackages/render"
	"github.com/utkangl/GoWEB/pkg"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	err := run()
	pkg.ErrorNilCheckPrint(err)

	fmt.Println("Starting the Application on Port: ", portNumber)

	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app)}

	err = serve.ListenAndServe()
	log.Fatal(err, "Cannot serve")

}

func run() error {

	// what we are going to put in session,
	gob.Register(models.Reservation{})

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // uses https insted of http

	app.Session = session

	tempCache, err := render.CreateTemplateCache()
	pkg.ErrorNilCheckFatal(err)
	pkg.ErrorNilCheckReturn(err)
	app.TemplateCache = tempCache

	repository := handlers.CreateRepo(&app)
	handlers.SetRepo(repository)

	render.SetConfig(&app)

	return nil
}
