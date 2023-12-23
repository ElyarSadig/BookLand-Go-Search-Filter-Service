package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	router.HandleFunc("/api/search/books/best-sellers", app.SearchBestSellerBooksHandler).Methods("GET")
	router.HandleFunc("/api/search/books/language-original", app.SearchOriginalLanguageBooksHandler).Methods("GET")
	router.HandleFunc("/api/search/books/popular", app.SearchPopularBooksHandler).Methods("GET")
	router.HandleFunc("/api/search/books/economical", app.SearchEconomicalBooksHandler).Methods("GET")
	router.HandleFunc("/api/search/books/free", app.SearchFreeBooksHandler).Methods("GET")
	router.HandleFunc("/api/search/books/recently-added", app.SearchRecentlyAddedBooks).Methods("GET")


	return router
}
