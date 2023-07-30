package main

import (
	"myapp/pkg/config"
	"myapp/pkg/handlers"
	"net/http"

	"github.com/gorilla/pat"
)

func routes(app *config.AppConfig) http.Handler {
	router := pat.New()

	router.Get("/", handlers.Repo.Home)
	router.Get("/about", handlers.Repo.About)

	return router
}
