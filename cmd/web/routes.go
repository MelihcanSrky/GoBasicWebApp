package main

import (
	"net/http"

	"github.com/MelihcanSrky/BasicWebApp/pkg/config"
	"github.com/MelihcanSrky/BasicWebApp/pkg/handlers"
	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
