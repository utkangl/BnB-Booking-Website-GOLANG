package config

import "html/template"

// this struct holds the application config
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
