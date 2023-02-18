package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// this struct holds the application config, config contains the template cache
// so we dont need to load the cache every time we visit the template
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
