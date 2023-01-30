package funcs

import (
	"fmt"
	"html/template"
	"log"
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

var tmplCache = make(map[string]*template.Template) // this map will store the datas thar parsedTemplate function throws (and it's type is *template.Teamplate)
// and the key to look up  to data is going to be a string

func RenderTmpl(Res http.ResponseWriter, t string) {

	var tmpl *template.Template
	var err error

	_, inMap := tmplCache[t] // return true if we already have the template in the cache

	if !inMap { // if inMap is false
		log.Println("New template hes been created")
		err = createTemplate(t) // creating the template with the string that we take as parameter of the function, and assigning variable err to possible returning error
		if err != nil {
			log.Println(err)
		}

	} else {
		log.Println("This template is already in the cache and on use")
	}

	tmpl = tmplCache[t] // t is the name of the template, we are assigning the template to a variable called tmpl
	err = tmpl.Execute(Res, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplate(t string) error {

	templates := []string{
		fmt.Sprintf("./templates/%s", t), // assigning the parameter t as the name of the template
		"./templates/base.layout.tmpl",   // setting the layout
	}

	tmpl, err := template.ParseFiles(templates...) // sending the string in the slice of string called templates to function ParseFiles,
	// and assigning returns template and error as tmpl and err

	if err != nil {
		return err // return error if there is any
	}

	tmplCache[t] = tmpl //add template to our cache of templates
	return nil
}
