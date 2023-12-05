package data

import (
	"database/sql"
	"time"
)

const dbTimeout = time.Second * 2

var db *sql.DB

func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		Book: Book{},
	}
}

type Models struct {
	Book Book
}

type Book struct {
	ID             int     `json:"id"`
	Publisher      *string `json:"publisher"`
	BookName       string  `json:"bookname"`
	AuthorName     string  `json:"authorname"`
	TranslatorName *string `json:"translatorname"`
	ReleasedDate   int     `json:"releaseddate"`
	BookCoverImage string  `json:"bookcoverimage"`
	Price          float64 `json:"price"`
	Description    *string `json:"description"`
	NumberOfPages  int     `json:"numberofpages"`
	Language       string  `json:"language"`
}

func (b *Book) GetBooksByName(db *sql.DB, bookName string) ([]Book, error) {
	query := `
        SELECT
            b.ID,
            u.PublicationsName AS Publisher,
            b.BookName,
            b.AuthorName,
            b.TranslatorName,
            b.ReleasedDate,
            b.BookCoverImage,
            b.Price,
            b.Description,
            b.NumberOfPages,
            l.Name AS Language
        FROM
            Books b
            JOIN Users u ON b.UserID = u.ID
            JOIN Languages l ON b.LanguageID = l.ID
        WHERE
            b.BookName LIKE $1 OR b.BookName = $2 AND b.IsDelete = FALSE
    `
	rows, err := db.Query(query, "%"+bookName+"%", bookName)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []Book

	for rows.Next() {
		var book Book
		err := rows.Scan(
			&book.ID,
			&book.Publisher,
			&book.BookName,
			&book.AuthorName,
			&book.TranslatorName,
			&book.ReleasedDate,
			&book.BookCoverImage,
			&book.Price,
			&book.Description,
			&book.NumberOfPages,
			&book.Language,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (b *Book) GetBooksByAuthor(db *sql.DB, authorName string) ([]Book, error) {
	query := `
        SELECT
            b.ID,
            u.PublicationsName AS Publisher,
            b.BookName,
            b.AuthorName,
            b.TranslatorName,
            b.ReleasedDate,
            b.BookCoverImage,
            b.Price,
            b.Description,
            b.NumberOfPages,
            l.Name AS Language
        FROM
            Books b
            JOIN Users u ON b.UserID = u.ID
            JOIN Languages l ON b.LanguageID = l.ID
        WHERE
            b.AuthorName LIKE $1 OR b.AuthorName = $2 AND b.IsDelete = FALSE
    `
	rows, err := db.Query(query, "%"+authorName+"%", authorName)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []Book

	for rows.Next() {
		var book Book
		err := rows.Scan(
			&book.ID,
			&book.Publisher,
			&book.BookName,
			&book.AuthorName,
			&book.TranslatorName,
			&book.ReleasedDate,
			&book.BookCoverImage,
			&book.Price,
			&book.Description,
			&book.NumberOfPages,
			&book.Language,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (b *Book) GetBooksByCategory(db *sql.DB, category string) ([]Book, error) {
	query := `
		SELECT
			b.id,
			u.PublicationsName AS Publisher,
			b.bookname,
			b.authorname,
			b.translatorname,
			b.releaseddate,
			b.bookcoverimage,
			b.price,
			b.description,
			b.numberofpages,
			l.name AS language
		FROM Books b
		JOIN
			BookCategories bc ON b.id = bc.bookid
		JOIN
			Categories c ON bc.categoryid = c.id
		JOIN
			Languages l ON b.languageid = l.id
		JOIN 
			Users u ON b.UserID = u.ID
		WHERE
			c.name = $1
			AND b.isdelete = FALSE
	`

	rows, err := db.Query(query, category)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []Book

	for rows.Next() {
		var book Book
		err := rows.Scan(
			&book.ID,
			&book.Publisher,
			&book.BookName,
			&book.AuthorName,
			&book.TranslatorName,
			&book.ReleasedDate,
			&book.BookCoverImage,
			&book.Price,
			&book.Description,
			&book.NumberOfPages,
			&book.Language,
		)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
