package data

type Book struct {
	ID             int     `json:"id"`
	Publisher      *string `json:"publisher"`
	BookName       string  `json:"bookname"`
	AuthorName     string  `json:"authorname"`
	TranslatorName *string `json:"translatorname"`
	ReleasedDate   int     `json:"releaseddate"`
	BookCoverImage string  `json:"bookcoverimage"`
	Price          float64 `json:"price"`
	NumberOfPages  int     `json:"numberofpages"`
	Language       string  `json:"language"`
}