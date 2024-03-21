package domain

import (
	"log"
	"scrapper/pkg/scrapper/infrastructure"
	"strings"

	"github.com/antchfx/htmlquery"
)

func GetArtist(url string) []string {

	b, StatusCode, err := infrastructure.GetData(url)
	if err != nil {
		return nil
	}
	doc, err := htmlquery.Parse(strings.NewReader(string(b)))
	if err != nil {
		return nil
	}

	if StatusCode == 404 {
		return []string{}
	}

	list := htmlquery.Find(doc, "//div[@class='songList-table']")
	log.Println("URL", url)

	URL := []string{}
	for _, n := range list {

		list2 := htmlquery.Find(n, "//a[@href]")

		for _, nn := range list2 {
			URL = append(URL, htmlquery.SelectAttr(nn, "href"))
		}
	}
	return URL
}
