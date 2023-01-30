package funcs

import (
	"fmt"
	"html/template"
	"net/http"
)

// homePage's handler function
func HomePage(Res http.ResponseWriter, Req *http.Request) {
	RenderTmpl(Res, "homePage.tmpl")
}

// aboutPage's handler function
func AboutPage(Res http.ResponseWriter, Req *http.Request) {
	RenderTmpl(Res, "aboutPage.tmpl")
}

func RenderTmpl(Res http.ResponseWriter, tmpl string) {

	parsedTemplate, _ := template.ParseFiles("./Templates/" + tmpl)
	err := parsedTemplate.Execute(Res, nil)

	if err != nil {
		fmt.Println("Error occured while parsing template", err)
		return
	}
}
