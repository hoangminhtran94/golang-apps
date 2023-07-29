package handlers

import (
	"myapp/pkg/config"
	"myapp/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
func (m *Repository) Home(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "home.html")
}

// About page handler
func (m *Repository) About(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "about.html")
}
