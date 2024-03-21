package domain

import (
	"log"
	"scrapper/pkg/scrapper/infrastructure"
	"strings"

	"github.com/antchfx/htmlquery"
)

type Song struct {
	Url    string
	Artist string
	Title  string
	Lyric  string
}

func Get(url string) *Song {

	b, StatusCode, err := infrastructure.GetData(url)
	if err != nil {
		return nil
	}
	doc, err := htmlquery.Parse(strings.NewReader(string(b)))
	if err != nil {
		return nil
	}

	if StatusCode == 404 {
		return nil
	}

	list := htmlquery.Find(doc, "//div[@class='page-container']")
	log.Println("URL", url)

	s := Song{}
	s.Url = url
	for _, n := range list {
		s.Lyric += htmlquery.InnerText(n)
	}
	title := htmlquery.FindOne(doc, "//h1/a")
	artist := htmlquery.FindOne(doc, "//h2/a")

	if title != nil {
		s.Title = htmlquery.InnerText(title)
	}
	if artist != nil {
		s.Artist = htmlquery.InnerText(artist)
	}

	if s.Title == "" && s.Artist == "" {
		return nil
	}

	return &s
}
