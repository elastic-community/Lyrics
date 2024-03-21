package main

import (
	input "scrapper/pkg/input/application"
	output "scrapper/pkg/output/application"
	scrapper "scrapper/pkg/scrapper/application"
	"scrapper/pkg/scrapper/domain"
)

func main() {
	lyrics := input.GetAll("Links.csv")
	songs := []*domain.Song{}
	for _, link := range lyrics {
		s := scrapper.Get(link.Link)
		if s == nil {
			continue
		}
		songs = append(songs, s...)
	}

	output.Save(songs)
}
