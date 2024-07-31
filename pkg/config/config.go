package config

import "html/template"

// holds the application config
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
