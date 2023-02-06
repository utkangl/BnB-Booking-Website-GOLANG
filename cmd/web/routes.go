package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/utkangl/GoWEB/internalPackages/config"
	"github.com/utkangl/GoWEB/internalPackages/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer) // Recovers from possible panics and gives detaield information about what went wrong
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/book", handlers.Repo.GetBook)
	mux.Post("/book", handlers.Repo.PostBook)

	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/kings_suit", handlers.Repo.Kings_suit)
	mux.Get("/regular_room", handlers.Repo.Regular_room)

	mux.Post("/book-json", handlers.Repo.AvailabilityJSON)

	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
