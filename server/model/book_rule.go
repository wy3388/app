package model

type BookRule struct {
	ID               uint   `json:"id"`
	SourceID         uint   `json:"source_id"`
	Search           string `json:"search"`
	SearchBookName   string `json:"search_book_name"`
	SearchAuthor     string `json:"search_author"`
	SearchNewChapter string `json:"search_new_chapter"`
	BookUrl          string `json:"book_url"`
	ImageUrl         string `json:"image_url"`
	BookName         string `json:"book_name"`
	Author           string `json:"author"`
	Info             string `json:"info"`
	UpdateTime       string `json:"update_time"`
	Chapter          string `json:"chapter"`
	ChapterName      string `json:"chapter_name"`
	ChapterUrl       string `json:"chapter_url"`
}
