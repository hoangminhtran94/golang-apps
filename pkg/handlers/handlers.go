package handlers

import (
	"myapp/pkg/render"
	"net/http"
)

func Home(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "home.html")
}

// About page handler
func About(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "about.html")
}
