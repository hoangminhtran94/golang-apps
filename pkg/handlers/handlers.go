package handlers

import (
	"myapp/pkg/config"
	"myapp/pkg/models"
	"myapp/pkg/render"
	"net/http"
)

// TemplateData hold page data

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
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	render.RenderTemplate(res, "home.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About page handler
func (m *Repository) About(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "about.html", &models.TemplateData{})
}
