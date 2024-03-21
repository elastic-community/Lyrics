package application

import (
	"scrapper/pkg/scrapper/domain"
)

func Get(url string) []*domain.Song {

	songs := []*domain.Song{}

	for _, link := range domain.GetGenere(url) {
		if link == "" {
			continue
		}
		song := domain.GetArtist("https://www.letras.mus.br" + link)

		for _, s := range song {
			if s == "" {
				continue
			}
			uri := domain.GetSong("https://www.letras.mus.br" + s)
			if uri == "" {
				continue
			}

			data := domain.Get("https://www.letras.mus.br" + uri)
			if data == nil {
				continue
			}

			songs = append(songs, data)
		}
	}

	return songs
}
