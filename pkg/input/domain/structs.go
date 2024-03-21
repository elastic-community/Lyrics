package domain

import (
	"encoding/csv"
	"log"
	"scrapper/pkg/input/infraestructure"
)

type Data struct {
	Link string
}

func GetAll(filePath string) []Data {
	f, err := infraestructure.Read(filePath)
	if err != nil {
		log.Panicln(err)
	}
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	links := []Data{}
	for _, link := range data {
		d := Data{}
		d.Link = link[0]
		links = append(links, d)
	}
	return links
}
