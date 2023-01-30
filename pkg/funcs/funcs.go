package funcs

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

// homePage's handler function
func HomePage(Res http.ResponseWriter, Req *http.Request) {
	RenderTemplate(Res, "home.page.tmpl")
}

// aboutPage's handler function
func AboutPage(Res http.ResponseWriter, Req *http.Request) {
	RenderTemplate(Res, "about.page.tmpl")
}

func RenderTemplate(Res http.ResponseWriter, tmpl string) {
	// create a template cache
	tempCache, err := createTemplateCache()
	if err != nil {
		log.Println(err)
	}

	// get requested template from cache
	temp, isExist := tempCache[tmpl]
	if !isExist {
		log.Fatal(err, "The template that this handler function tries to pass as argument does not exist") // kill the program if template does not exist
	}

	buf := new(bytes.Buffer)     //Using buffer for higher protection,
	err = temp.Execute(buf, nil) //Rather than executing the template directly, it will executes its bytes.
	if err != nil {              // It will help to understand why does the error exactly come from
		log.Println(err)
	}
	//render the template
	_, err = buf.WriteTo(Res)
	if err != nil {
		log.Println(err)
	}
}

// this function takes all templates and layouts, puts them together and adds them to the cache
func createTemplateCache() (map[string]*template.Template, error) {
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

//var tmplCache = make(map[string]*template.Template) // this map will store the datas thar parsedTemplate function throws (and it's type is *template.Teamplate)
// and the key to look up  to data is going to be a string

// func RenderTemplate(Res http.ResponseWriter, t string) {

// 	var tmpl *template.Template
// 	var err error

// 	_, inMap := tmplCache[t] // return true if we already have the template in the cache

// 	if !inMap { // if inMap is false
// 		log.Println("New template hes been created")
// 		err = CreateTemplateCache(t) // creating the template with the string that we take as parameter of the function, and assigning variable err to possible returning error
// 		if err != nil {
// 			log.Println(err)
// 		}

// 	} else {
// 		log.Println("This template is already in the cache and on use")
// 	}

// 	tmpl = tmplCache[t] // t is the name of the template, we are assigning the template to a variable called tmpl
// 	err = tmpl.Execute(Res, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func CreateTemplateCache(t string) error {

// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t), // assigning the parameter t as the name of the template
// 		"./templates/base.layout.tmpl",   // setting the layout
// 	}

// 	tmpl, err := template.ParseFiles(templates...) // sending the string in the slice of string called templates to function ParseFiles,
// 	// and assigning returns template and error as tmpl and err

// 	if err != nil {
// 		return err // return error if there is any
// 	}

// 	tmplCache[t] = tmpl //add template to our cache of templates
// 	return nil
