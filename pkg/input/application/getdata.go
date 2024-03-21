package application

import (
	"scrapper/pkg/input/domain"
)

func GetAll(filepath string) []domain.Data {
	return domain.GetAll(filepath)
}
