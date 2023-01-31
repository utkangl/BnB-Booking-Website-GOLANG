package config

import "html/template"

// this struct holds the application config, config contains the template cache
// so we dont need to load the cache every time we visit the template
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
