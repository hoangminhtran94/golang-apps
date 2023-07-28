package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(res http.ResponseWriter, templ string) {

	parsedTemplate, err0 := template.ParseFiles("../../templates/"+templ, "../../templates/base.html")
	if err0 != nil {
		fmt.Println("Error parsing template:", err0)
		return
	}
	err := parsedTemplate.Execute(res, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}
