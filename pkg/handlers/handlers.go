package handlers

import (
	"net/http"

	"github.com/utkangl/GoWEB/pkg/config"
	"github.com/utkangl/GoWEB/pkg/render"
)

// Repository variable to used by handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func CreateRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets the repository for the handlers
func SetRepo(r *Repository) {
	Repo = r
}

// homePage's handler function (repository typed receiver function )
func (rep *Repository) HomePage(Res http.ResponseWriter, Req *http.Request) {
	render.RenderTemplate(Res, "home.page.tmpl")
}

// aboutPage's handler function (repository typed receiver function )
func (rep *Repository) AboutPage(Res http.ResponseWriter, Req *http.Request) {
	render.RenderTemplate(Res, "about.page.tmpl")
}
