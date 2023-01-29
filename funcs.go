package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomePage(Res http.ResponseWriter, Req *http.Request) {
	renderTmpl(Res, "homePage.tmpl")
}

// this function is a handler function
func AboutPage(Res http.ResponseWriter, Req *http.Request) {
	renderTmpl(Res, "aboutPage.tmpl")
}

func renderTmpl(Res http.ResponseWriter, tmpl string) {

	parsedTemplate, _ := template.ParseFiles("./Templates/" + tmpl)
	err := parsedTemplate.Execute(Res, nil)

	if err != nil {
		fmt.Println("Error occured while parsing template", err)
		return
	}
}
