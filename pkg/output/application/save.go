package application

import (
	"encoding/csv"
	"io"
	"os"
	"scrapper/pkg/scrapper/domain"

	"github.com/gocarina/gocsv"
)

func Save(data []*domain.Song) {

	file, err := os.OpenFile("output.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = '|'
		return r // Allows use pipe as delimiter
	})

	gocsv.MarshalFile(&data, file)

}
