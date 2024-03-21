package domain

import (
	"scrapper/pkg/scrapper/infrastructure"
	"strings"

	"github.com/antchfx/htmlquery"
)

func GetSong(url string) string {

	b, StatusCode, err := infrastructure.GetData(url)
	if err != nil {
		return ""
	}

	if StatusCode == 404 {
		return ""
	}
	doc, err := htmlquery.Parse(strings.NewReader(string(b)))
	if err != nil {
		return ""
	}

	data := htmlquery.FindOne(doc, "//a[@class='js-lyric-event lyric-menu-print']")

	URL := htmlquery.SelectAttr(data, "href")

	return URL
}
