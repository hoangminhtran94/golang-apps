package main

import (
	"fmt"
	"myapp/pkg/handlers"
	"net/http"
)

const portNumber = ":3080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println("Starting server at" + portNumber)
	http.ListenAndServe(portNumber, nil)
}
