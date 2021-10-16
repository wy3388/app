package model

type BookSelf struct {
	ID         uint   `json:"id"`
	BookName   string `json:"book_name"`
	BookUrl    string `json:"book_url"`
	ChapterUrl string `json:"chapter_url"`
}
