package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/larnTechGeeks/basis-web/pkg/config"
	"github.com/larnTechGeeks/basis-web/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// Middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	// routes
	mux.Get("/", handlers.Repo.HomePage)
	mux.Get("/about", handlers.Repo.AboutPage)

	return mux
}
