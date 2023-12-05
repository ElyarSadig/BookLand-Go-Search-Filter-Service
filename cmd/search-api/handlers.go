package main

import (
	"net/http"
)

func (app *Config) SearchBooksByNameHandler(w http.ResponseWriter, r *http.Request) {
	bookName := r.URL.Query().Get("bookName")

	books, err := app.Models.Book.GetBooksByName(app.DB, bookName)

	if err != nil {
		app.errorJSON(w, err, "خطایی در سرور رخ داده است", "InternalServerError")
		return
	}

	if len(books) == 0 {
		app.writeJSON(w, http.StatusNotFound, "")
		return
	}

	app.writeJSON(w, http.StatusOK, books)
}

func (app *Config) SearchBooksByAuthorHandler(w http.ResponseWriter, r *http.Request) {
	authorName := r.URL.Query().Get("author")

	books, err := app.Models.Book.GetBooksByAuthor(app.DB, authorName)

	if err != nil {
		app.errorJSON(w, err, "خطایی در سرور رخ داده است", "InternalServerError")
		return
	}

	if len(books) == 0 {
		app.writeJSON(w, http.StatusNotFound, "")
		return
	}

	app.writeJSON(w, http.StatusOK, books)
}

func (app *Config) SearchBooksByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	categoryName := r.URL.Query().Get("category")

	books, err := app.Models.Book.GetBooksByCategory(app.DB, categoryName)

	if err != nil {
		app.errorJSON(w, err, "خطایی در سرور رخ داده است", "InternalServerError")
		return
	}

	if len(books) == 0 {
		app.writeJSON(w, http.StatusNotFound, "")
		return
	}

	app.writeJSON(w, http.StatusOK, books)
}
