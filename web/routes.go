package main

import (
	"net/http"
	"udemyCourse1/pkg/config"
	"udemyCourse1/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/generals-quarters", http.HandlerFunc(handlers.Repo.Generals))
	mux.Get("/majors-suite", http.HandlerFunc(handlers.Repo.Majors))
	mux.Get("/search-availability", http.HandlerFunc(handlers.Repo.Availability))
	mux.Get("/contact", http.HandlerFunc(handlers.Repo.Contacts))
	
	mux.Get("/make-reservation", handlers.Repo.Reservation)
	
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
