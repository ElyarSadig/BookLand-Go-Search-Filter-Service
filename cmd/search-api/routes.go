package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *Config) routes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/api/search/books", app.SearchBooksByNameHandler).Methods("GET").Queries("bookName", "{bookName}")
	router.HandleFunc("/api/search/books", app.SearchBooksByAuthorHandler).Methods("GET").Queries("author", "{author}")
	router.HandleFunc("/api/search/books", app.SearchBooksByCategoryHandler).Methods("GET").Queries("category", "{category}")

	return router
}
