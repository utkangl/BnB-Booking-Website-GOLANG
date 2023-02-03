package handlers

import (
	"encoding/json"
	"fmt"
	"log"
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

type jsonRes struct {
	OK      bool   `json:"ok"`
	Message string `json:"msg"`
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

func (rep *Repository) Home(Res http.ResponseWriter, Req *http.Request) {

	stringMap := make(map[string]string)
	stringMap["greeting"] = "Welcome to HomePage!"

	VisiterIP := Req.RemoteAddr
	rep.App.Session.Put(Req.Context(), "VisiterIP", VisiterIP)

	render.RenderTemplate(Res, "home.page.tmpl", &models.TemplateData{StringMap: stringMap}, Req)
}

func (rep *Repository) About(Res http.ResponseWriter, Req *http.Request) {

	stringMap := make(map[string]string)
	stringMap["greeting"] = "Welcome to AboutPage!"

	VisiterIP := rep.App.Session.GetString(Req.Context(), "VisiterIP")
	stringMap["VisiterIP"] = VisiterIP

	render.RenderTemplate(Res, "about.page.tmpl", &models.TemplateData{StringMap: stringMap}, Req)
}

func (rep *Repository) Contact(Res http.ResponseWriter, Req *http.Request) {
	render.RenderTemplate(Res, "contact.page.tmpl", &models.TemplateData{}, Req)
}

func (rep *Repository) GetBook(Res http.ResponseWriter, Req *http.Request) {
	render.RenderTemplate(Res, "book.page.tmpl", &models.TemplateData{}, Req)
}

func (rep *Repository) PostBook(Res http.ResponseWriter, Req *http.Request) {

	start := Req.Form.Get("start")
	end := Req.Form.Get("end")

	Res.Write([]byte(fmt.Sprintf("start and finish date are %s and %s", start, end)))

}

func (rep *Repository) AvailabilityJSON(Res http.ResponseWriter, Req *http.Request) {

	jsonResponse := jsonRes{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(jsonResponse, "", "   ")
	if err != nil {
		log.Println(err)
	}

	Res.Header().Set("Content-Type", "application/json")
	Res.Write(out)

}

func (rep *Repository) Make_reservation(Res http.ResponseWriter, Req *http.Request) {
	render.RenderTemplate(Res, "make-reservation.page.tmpl", &models.TemplateData{}, Req)
}

func (rep *Repository) Kings_suit(Res http.ResponseWriter, Req *http.Request) {
	render.RenderTemplate(Res, "kings_suit.page.tmpl", &models.TemplateData{}, Req)
}

func (rep *Repository) Regular_room(Res http.ResponseWriter, Req *http.Request) {
	render.RenderTemplate(Res, "regular-room.page.tmpl", &models.TemplateData{}, Req)
}
