package service

import (
	"app/server/common/parse"
	"app/server/model"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

type SearchReq struct {
	Key string `json:"key"`
	Id  string `json:"id"`
}

type Search struct {
	BookName   string `json:"book_name"`
	Author     string `json:"author"`
	NewChapter string `json:"new_chapter"`
	BookUrl    string `json:"book_url"`
}

type InfoReq struct {
	Url string `json:"url"`
	Id  string `json:"id"`
}

type Book struct {
	BookName    string     `json:"book_name"`
	Author      string     `json:"author"`
	UpdateTime  string     `json:"update_time"`
	Info        string     `json:"info"`
	ImageUrl    string     `json:"image_url"`
	ChapterList []*Chapter `json:"chapter"`
}

type Chapter struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

type BodyReq struct {
	Url string `json:"url"`
	Id  string `json:"id"`
}

type Content struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type BookService struct {
}

var p parse.Parse

func (BookService) Source() (s []*model.Source) {
	if err := db.Find(&s).Error; err != nil {
		panic(err)
	}
	return
}

func (BookService) Search(req *SearchReq) (sl []*Search) {
	var s model.Source
	if err := db.Preload("BookRule").First(&s, req.Id).Error; err != nil {
		panic(err)
	}
	u := fmt.Sprintf(s.SearchUrl, req.Key)
	p.Url(u, s).List(s.BookRule.Search).Each(func(i int, selection *goquery.Selection) {
		bookName := p.Selection(selection, s).First(s.BookRule.SearchBookName).Text()
		author := p.Selection(selection, s).First(s.BookRule.SearchAuthor).Text()
		newChapter := p.Selection(selection, s).First(s.BookRule.SearchNewChapter).Text()
		bookUrl := p.Selection(selection, s).First(s.BookRule.BookUrl).Attr()
		sl = append(sl, &Search{
			BookName:   bookName,
			Author:     author,
			NewChapter: newChapter,
			BookUrl:    bookUrl,
		})
	})
	return
}

func (BookService) Info(req *InfoReq) *Book {
	var s model.Source
	if err := db.Preload("BookRule").First(&s, req.Id).Error; err != nil {
		panic(err)
	}
	doc := p.Url(req.Url, s)
	bookName := doc.First(s.BookRule.BookName).Text()
	author := doc.First(s.BookRule.Author).Text()
	updateTime := doc.First(s.BookRule.UpdateTime).Text()
	info := doc.First(s.BookRule.Info).Text()
	imageUrl := doc.First(s.BookRule.ImageUrl).Attr()
	var chapterList []*Chapter
	doc.List(s.BookRule.Chapter).Each(func(i int, selection *goquery.Selection) {
		chapterName := p.Selection(selection, s).First(s.BookRule.ChapterName).Text()
		chapterUrl := p.Selection(selection, s).First(s.BookRule.ChapterUrl).Attr()
		chapterList = append(chapterList, &Chapter{
			Title: chapterName,
			Url:   chapterUrl,
		})
	})
	return &Book{
		BookName:    bookName,
		Author:      author,
		UpdateTime:  updateTime,
		Info:        info,
		ImageUrl:    imageUrl,
		ChapterList: chapterList,
	}
}

func (BookService) Body(req *BodyReq) *Content {
	var s model.Source
	if err := db.Preload("BookRule").First(&s, req.Id).Error; err != nil {
		panic(err)
	}
	doc := p.Url(req.Url, s)
	title := doc.First(s.BookRule.ContentTitle).Text()
	body := doc.First(s.BookRule.ContentBody).Text()
	return &Content{
		Title: title,
		Body:  body,
	}
}
