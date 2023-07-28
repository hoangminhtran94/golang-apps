package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":3080"

// Home page handler
func renderTemplate(res http.ResponseWriter, templ string) {

	parsedTemplate, _ := template.ParseFiles("./templates/" + templ)
	err := parsedTemplate.Execute(res, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}

func Home(res http.ResponseWriter, req *http.Request) {
	renderTemplate(res, "home.page.tmpl")
}

// About page handler
func About(res http.ResponseWriter, req *http.Request) {

}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.ListenAndServe(portNumber, nil)
}
