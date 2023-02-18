package handlers

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/justinas/nosurf"
	"github.com/utkangl/GoWEB/internalPackages/config"
	"github.com/utkangl/GoWEB/internalPackages/models"
	"github.com/utkangl/GoWEB/internalPackages/render"
	"github.com/utkangl/GoWEB/pkg"
)

var app config.AppConfig
var session *scs.SessionManager
var pathToTemplates = "./../../templates"

func getRoutes() http.Handler {

	gob.Register(models.Reservation{})

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // uses https insted of http

	app.Session = session

	tempCache, err := CreateTestTemplateCache()
	pkg.ErrorNilCheckFatal(err)
	//pkg.ErrorNilCheckReturn(err)
	app.TemplateCache = tempCache
	app.UseCache = true

	repository := CreateRepo(&app)
	SetRepo(repository)

	render.SetConfig(&app)

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer) // Recovers from possible panics and gives detaield information about what went wrong
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/book", Repo.GetBook)
	mux.Post("/book", Repo.PostBook)

	mux.Get("/contact", Repo.Contact)

	mux.Get("/kings_suit", Repo.Kings_suit)
	mux.Get("/regular_room", Repo.Regular_room)

	mux.Post("/book-json", Repo.AvailabilityJSON)

	mux.Get("/make-reservation", Repo.Reservation)
	mux.Post("/make-reservation", Repo.PostReservation)
	mux.Get("/reservation-summary", Repo.ReservationSummary)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
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

func CreateTestTemplateCache() (map[string]*template.Template, error) {
	templateCache := make(map[string]*template.Template)

	// get all templates from templates folder
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return templateCache, err
	}

	//assign the name of template and parse the template
	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)

		if err != nil {
			return templateCache, err
		}

		// get all layouts
		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))

		if err != nil {
			return templateCache, err
		}

		// add templates and layouts up with parseglob function
		if len(matches) > 0 {
			templateSet, err := templateSet.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return templateCache, err
			}

			// add the final template to cache
			templateCache[name] = templateSet
		}
	}
	return templateCache, nil
}
