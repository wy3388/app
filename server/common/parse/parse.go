package parse

import (
	"app/server/common"
	"app/server/model"
	"app/server/utils"
	"app/server/utils/req"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
)

type rule struct {
	Base    bool      `json:"base"`
	Attr    string    `json:"attr"`
	Replace []replace `json:"replace"`
	Regexp  string    `json:"regexp"`
	Type    string    `json:"type"`
}

type replace struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type parse struct {
	doc    *goquery.Document
	sel    *goquery.Selection
	isDoc  bool
	source model.Source
	rule   rule
}

type Parse struct {
}

func (Parse) Document(doc *goquery.Document, source model.Source) *parse {
	return &parse{
		doc:    doc,
		source: source,
		isDoc:  true,
	}
}

func (Parse) Selection(sel *goquery.Selection, source model.Source) *parse {
	return &parse{
		sel:    sel,
		source: source,
		isDoc:  false,
	}
}

func (Parse) Url(url string, source model.Source) *parse {
	doc, err := req.Document(url, source)
	if err != nil {
		panic(err)
	}
	return &parse{
		doc:    doc,
		source: source,
		isDoc:  true,
	}
}

func (p *parse) First(query string) *parse {
	if strings.Contains(query, common.RuleSymbol) {
		sp := strings.Split(query, common.RuleSymbol)
		if sp[0] != "" {
			if p.isDoc {
				p.sel = p.doc.Find(sp[0]).First()
			} else {
				p.sel = p.sel.Find(sp[0]).First()
			}
		} else {
			if p.isDoc {
				p.sel = p.doc.Selection
			}
		}
		utils.ToObj(sp[len(sp)-1], &p.rule)
	} else {
		if query != "" {
			if p.isDoc {
				p.sel = p.doc.Find(query).First()
			} else {
				p.sel = p.sel.Find(query).First()
			}
		} else {
			if p.isDoc {
				p.sel = p.doc.Selection
			}
		}
	}
	return p
}

func (p *parse) List(query string) *goquery.Selection {
	if p.isDoc {
		return p.doc.Find(query)
	}
	return p.sel.Find(query)
}

func (p *parse) Text() string {
	if p.sel == nil {
		return ""
	}
	text := ""
	if p.rule.Type == "" || p.rule.Type == "text" {
		text = p.sel.Text()
	} else {
		text = p.Html()
	}
	if p.rule.Regexp != "" {
		text = regexp.MustCompile(p.rule.Regexp).FindString(text)
	}
	if len(p.rule.Replace) > 0 {
		for _, r := range p.rule.Replace {
			text = strings.ReplaceAll(text, r.Key, r.Value)
		}
	}
	return text
}

func (p *parse) Html() string {
	html, err := p.sel.Html()
	if err != nil {
		return ""
	}
	return html
}

func (p *parse) Attr() string {
	if p.sel == nil || p.rule.Attr == "" {
		return ""
	}
	attr := p.sel.AttrOr(p.rule.Attr, "")
	if p.rule.Base {
		return p.source.Url + attr
	}
	return attr
}
