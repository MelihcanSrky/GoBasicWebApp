package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MelihcanSrky/BasicWebApp/pkg/config"
	"github.com/MelihcanSrky/BasicWebApp/pkg/handlers"
	"github.com/MelihcanSrky/BasicWebApp/pkg/render"
)

const port = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting app on port %s", port))
	_ = http.ListenAndServe(port, nil)
}
