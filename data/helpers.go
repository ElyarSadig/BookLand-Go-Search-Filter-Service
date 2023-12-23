package data

import "database/sql"

func getBooksByQuery(db *sql.DB, query string, args ...interface{}) ([]Book, error) {
	rows, err := db.Query(query, args...)

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
