package data

import (
	"database/sql"
)

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

func (b *Book) GetBooksByName(db *sql.DB, bookName string) ([]Book, error) {
	query := `
        SELECT
            b.ID,
            u.PublicationsName AS publisher,
            b.BookName,
            b.AuthorName,
            b.TranslatorName,
            b.ReleasedDate,
            b.BookCoverImage,
            b.Price,
            b.NumberOfPages,
            l.Name AS Language
        FROM
            Books b
            JOIN Users u ON b.UserID = u.ID
            JOIN Languages l ON b.LanguageID = l.ID
        WHERE
            b.BookName LIKE $1 OR b.BookName = $2 AND b.IsDelete = FALSE
    `

	books, err := getBooksByQuery(db, query, "%"+bookName+"%", bookName)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (b *Book) GetBooksByAuthor(db *sql.DB, authorName string) ([]Book, error) {
	query := `
        SELECT
            b.ID,
            u.PublicationsName AS publisher,
            b.BookName,
            b.AuthorName,
            b.TranslatorName,
            b.ReleasedDate,
            b.BookCoverImage,
            b.Price,
            b.NumberOfPages,
            l.Name AS Language
        FROM
            Books b
            JOIN Users u ON b.UserID = u.ID
            JOIN Languages l ON b.LanguageID = l.ID
        WHERE
            b.AuthorName LIKE $1 OR b.AuthorName = $2 AND b.IsDelete = FALSE
    `

	books, err := getBooksByQuery(db, query, "%"+authorName+"%", authorName)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (b *Book) GetBooksByCategory(db *sql.DB, category string) ([]Book, error) {
	query := `
		SELECT
			b.id,
			u.PublicationsName AS publisher,
			b.bookname,
			b.authorname,
			b.translatorname,
			b.releaseddate,
			b.bookcoverimage,
			b.price,
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
			c.name = $1 AND b.isdelete = FALSE
	`

	books, err := getBooksByQuery(db, query, category)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (b *Book) GetBestSellerBooks(db *sql.DB) ([]Book, error) {
	query := `
		SELECT
			b.Id,
			b.BookName,
			u.PublicationsName AS publisher,
			b.AuthorName,
			b.TranslatorName,
			b.ReleasedDate,
			b.BookCoverImage,
			b.Price,
			b.NumberOfPages,
			l.Name AS Language
		FROM
			UserBooks ub
		JOIN
			Books b ON ub.BookId = b.Id
		JOIN 
			Users u ON b.UserID = u.ID
		JOIN
			Languages l ON b.LanguageId = l.Id
		WHERE
			b.isdelete = FALSE
		GROUP BY
			b.Id, b.BookName, b.AuthorName, b.TranslatorName, b.ReleasedDate,
			b.BookCoverImage, b.Price, b.NumberOfPages, l.Name, publisher
		ORDER BY
			COUNT(ub.BookId) DESC;
	`

	books, err := getBooksByQuery(db, query)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (b *Book) GetOriginalLanguageBooks(db *sql.DB) ([]Book, error) {
	query := `
        SELECT
            b.ID,
            u.PublicationsName AS publisher,
            b.BookName,
            b.AuthorName,
            b.TranslatorName,
            b.ReleasedDate,
            b.BookCoverImage,
            b.Price,
            b.NumberOfPages,
            l.Name AS Language
        FROM Books b
        JOIN Users u ON b.UserID = u.ID
        JOIN Languages l ON b.LanguageID = l.ID
        WHERE
            b.LanguageID <> 1 AND b.IsDelete = FALSE
	`

	books, err := getBooksByQuery(db, query)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (b *Book) GetPopularBooks(db *sql.DB) ([]Book, error) {
	query := `
		SELECT
			b.Id,
			u.PublicationsName AS publisher,
			b.BookName,
			b.AuthorName,
			b.TranslatorName,
			b.ReleasedDate,
			b.BookCoverImage,
			b.Price,
			b.NumberOfPages,
			l.Name AS Language
		FROM
			Books b
		JOIN
			Languages l ON b.LanguageId = l.Id
		JOIN 
			Users u ON b.UserID = u.ID
		WHERE
			b.isdelete = FALSE
			AND b.Id IN (
				SELECT
					BookId
				FROM
					Reviews
				GROUP BY
					BookId
				HAVING
					AVG(Rating) >= 3.5   	-- Adjust the threshold as needed
			)
		ORDER BY
			(SELECT AVG(Rating) FROM Reviews WHERE BookId = b.Id) DESC;
	`

	books, err := getBooksByQuery(db, query)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (b *Book) GetEconomicalBooks(db *sql.DB) ([]Book, error) {
	query := `
		SELECT
			b.ID,
			u.PublicationsName AS publisher,
			b.BookName,
			b.AuthorName,
			b.TranslatorName,
			b.ReleasedDate,
			b.BookCoverImage,
			b.Price,
			b.NumberOfPages,
			l.Name AS Language
		FROM
			Books b
			JOIN Users u ON b.UserID = u.ID
			JOIN Languages l ON b.LanguageID = l.ID
		WHERE
			b.IsDelete = FALSE AND b.Price <> 0 AND b.Price < 60000
		ORDER BY
			b.Price ASC;
	`

	books, err := getBooksByQuery(db, query)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (b *Book) GetFreeBooks(db *sql.DB) ([]Book, error) {
	query := `
        SELECT
            b.ID,
            u.PublicationsName AS publisher,
            b.BookName,
            b.AuthorName,
            b.TranslatorName,
            b.ReleasedDate,
            b.BookCoverImage,
            b.Price,
            b.NumberOfPages,
            l.Name AS Language
        FROM
            Books b
        JOIN Users u ON b.UserID = u.ID
        JOIN Languages l ON b.LanguageID = l.ID
        WHERE
            b.IsDelete = FALSE AND b.Price = 0
    `

	books, err := getBooksByQuery(db, query)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (b *Book) GetRecentlyAddedBooks(db *sql.DB) ([]Book, error) {
	query := `
		SELECT
            b.ID,
            u.PublicationsName AS publisher,
            b.BookName,
            b.AuthorName,
            b.TranslatorName,
            b.ReleasedDate,
            b.BookCoverImage,
            b.Price,
            b.NumberOfPages,
            l.Name AS Language
        FROM
            Books b
        JOIN Users u ON b.UserID = u.ID
        JOIN Languages l ON b.LanguageID = l.ID
        WHERE
            b.IsDelete = FALSE
		ORDER BY
			b.ReleasedDate DESC;
	`

	books, err := getBooksByQuery(db, query)

	if err != nil {
		return nil, err
	}

	return books, nil
}
