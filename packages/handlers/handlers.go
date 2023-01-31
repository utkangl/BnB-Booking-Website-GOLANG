package handlers

import (
	"net/http"

	"github.com/utkangl/GoWEB/packages/config"
	"github.com/utkangl/GoWEB/packages/models"
	"github.com/utkangl/GoWEB/packages/render"
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

	stringMap := make(map[string]string)
	stringMap["greeting"] = "Welcome to HomePage!"

	VisiterIP := Req.RemoteAddr
	rep.App.Session.Put(Req.Context(), "VisiterIP", VisiterIP)

	render.RenderTemplate(Res, "home.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

// aboutPage's handler function (repository typed receiver function )
func (rep *Repository) AboutPage(Res http.ResponseWriter, Req *http.Request) {

	stringMap := make(map[string]string)
	stringMap["greeting"] = "Welcome to AboutPage!"

	VisiterIP := rep.App.Session.GetString(Req.Context(), "VisiterIP")
	stringMap["VisiterIP"] = VisiterIP

	render.RenderTemplate(Res, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
