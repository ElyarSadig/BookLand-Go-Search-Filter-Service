package main

import (
	"net/http"

	"github.com/rs/cors"
	"github.com/gorilla/mux"
)

func (app *Config) routes() http.Handler {
	router := mux.NewRouter()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://localhost:3000", "http://*"},
		AllowedMethods:   []string{"GET", "OPTIONS", "HEAD"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	// Apply CORS middleware to the router
	router.Use(corsHandler.Handler)

	router.HandleFunc("/api/search/books", app.SearchBooksByNameHandler).Methods("GET").Queries("bookName", "{bookName}")
	router.HandleFunc("/api/search/books", app.SearchBooksByAuthorHandler).Methods("GET").Queries("author", "{author}")
	router.HandleFunc("/api/search/books", app.SearchBooksByCategoryHandler).Methods("GET").Queries("category", "{category}")

	return router
}
