package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/utkangl/GoWEB/packages/config"
	"github.com/utkangl/GoWEB/packages/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.HomePage)
	mux.Get("/about", handlers.Repo.AboutPage)

	return mux
}
