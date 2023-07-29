package main

import (
	"fmt"
	"log"
	"myapp/pkg/config"
	"myapp/pkg/handlers"
	"myapp/pkg/render"
	"net/http"
)

const portNumber = ":3080"

func main() {
	var app config.AppConfig

	template, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = template
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println("Starting server at" + portNumber)
	http.ListenAndServe(portNumber, nil)
}
