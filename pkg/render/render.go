package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(res http.ResponseWriter, templ string) {
	//Create template caches
	caches, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	//getRequested template cache from caches
	cache, ok := caches[templ]
	if !ok {
		log.Fatal(err)
	}

	err = cache.Execute(res, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("../../templates/*.html")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		cache, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("../../templates/*.layout.html")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			cache, err = cache.ParseGlob("../../templates/*.layout.html")
		}
		if err != nil {
			return myCache, err
		}
		myCache[name] = cache
	}
	return myCache, nil
}
