package model

import "time"

type BookSelf struct {
	ID           uint      `json:"id"`
	BookName     string    `json:"book_name"`
	BookUrl      string    `json:"book_url"`
	ChapterUrl   string    `json:"chapter_url"`
	ChapterName  string    `json:"chapter_name"`
	ChapterTotal uint      `json:"chapter_total"`
	ChapterIndex uint      `json:"chapter_index"`
	SourceId     uint      `json:"source_id"`
	CreateTime   time.Time `json:"create_time"`
}
