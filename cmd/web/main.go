package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jimsyyap/bnb_go/pkg/config"
	"github.com/jimsyyap/bnb_go/pkg/handlers"
	"github.com/jimsyyap/bnb_go/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println("Starting application on port", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}
