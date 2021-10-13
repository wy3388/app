package model

type Source struct {
	ID        uint     `json:"id"`
	Name      string   `json:"name"`
	Url       string   `json:"url"`
	SearchUrl string   `json:"search_url"`
	BookRule  BookRule `json:"book_rule"`
}
