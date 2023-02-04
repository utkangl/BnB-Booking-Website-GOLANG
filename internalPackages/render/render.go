package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/utkangl/GoWEB/internalPackages/config"
	"github.com/utkangl/GoWEB/internalPackages/models"
)

var app *config.AppConfig

func SetConfig(Application *config.AppConfig) {
	app = Application
}

// we will call this function when we want some data to be sent to every template of our application
func AddDefaultDataToTemplate(tempData *models.TemplateData, req *http.Request) *models.TemplateData {
	tempData.CSRFToken = nosurf.Token(req)
	return tempData
}

func RenderTemplate(Res http.ResponseWriter, tmpl string, templateData *models.TemplateData, req *http.Request) {

	// create a template cache
	tempCache := app.TemplateCache

	// get requested template from cache
	temp, isExist := tempCache[tmpl]
	if !isExist {
		log.Fatal("The template that this handler function tries to pass as argument does not exist") // kill the program if template does not exist
	}

	buf := new(bytes.Buffer)

	templateData = AddDefaultDataToTemplate(templateData, req)

	//Using buffer for higher protection,
	err := temp.Execute(buf, templateData) //Rather than executing the template directly, it will executes its bytes.
	if err != nil {                        // It will help to understand why does the error exactly come from
		log.Println(err)
	}
	//render the template
	_, err = buf.WriteTo(Res)
	if err != nil {
		log.Println(err)
	}
}

// this function takes all templates and layouts, puts them together and adds them to the cache
func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := make(map[string]*template.Template)

	// get all templates from templates folder
	pages, err := filepath.Glob("./Templates/*.page.tmpl")
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
		matches, err := filepath.Glob("./Templates/*.layout.tmpl")

		if err != nil {
			return templateCache, err
		}

		// add templates and layouts up with parseglob function
		if len(matches) > 0 {
			templateSet, err := templateSet.ParseGlob("./Templates/*.layout.tmpl")
			if err != nil {
				return templateCache, err
			}

			// add the final template to cache
			templateCache[name] = templateSet
		}
	}
	return templateCache, nil
}
