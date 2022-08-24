package main

import (
	"github.com/code-chimp/gobookings/pkg/config"
	"github.com/code-chimp/gobookings/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewMux()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/generals-quarters", handlers.Repo.GeneralsQuarters)
	mux.Get("/colonels-suite", handlers.Repo.ColonelsSuite)
	mux.Get("/make-reservation", handlers.Repo.MakeReservation)
	mux.Get("/search-availability", handlers.Repo.SearchAvailability)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
