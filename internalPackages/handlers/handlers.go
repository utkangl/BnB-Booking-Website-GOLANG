package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/utkangl/GoWEB/internalPackages/config"
	"github.com/utkangl/GoWEB/internalPackages/forms"
	"github.com/utkangl/GoWEB/internalPackages/models"
	"github.com/utkangl/GoWEB/internalPackages/render"
	"github.com/utkangl/GoWEB/pkg"
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

func (rep *Repository) Reservation(Res http.ResponseWriter, Req *http.Request) {

	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.RenderTemplate(Res, "make-reservation.page.tmpl", &models.TemplateData{Form: forms.CreateForm(nil), Data: data}, Req)
}

func (rep *Repository) PostReservation(Res http.ResponseWriter, Req *http.Request) {
	err := Req.ParseForm()
	pkg.ErrorNilCheckPrint(err)

	reservation := models.Reservation{
		FirstName: Req.Form.Get("first_name"),
		LastName:  Req.Form.Get("last_name"),
		Phone:     Req.Form.Get("phone"),
		Email:     Req.Form.Get("email"),
	}

	// PostForm is url.Values type
	form := forms.CreateForm(Req.PostForm)

	// function has two parameters, field String and http.Request. Return true if the field is not empty
	//form.Has("first_name", Req)

	form.Required("first_name", "last_name", "phone", "email")
	form.MinLength("first_name", 5, Req)
	form.MinLength("last_name", 8, Req)
	form.MinLength("phone", 9, Req)

	form.IsValidEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(Res, "make-reservation.page.tmpl", &models.TemplateData{Form: form, Data: data}, Req)

		return // stop processing
	}
}

func (rep *Repository) Kings_suit(Res http.ResponseWriter, Req *http.Request) {
	render.RenderTemplate(Res, "kings_suit.page.tmpl", &models.TemplateData{}, Req)
}

func (rep *Repository) Regular_room(Res http.ResponseWriter, Req *http.Request) {
	render.RenderTemplate(Res, "regular-room.page.tmpl", &models.TemplateData{}, Req)
}
