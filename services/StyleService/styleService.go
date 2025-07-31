package styleservice

import (
	"github.com/gorgoroth31/boulder-tracker/models"
	stylerepository "github.com/gorgoroth31/boulder-tracker/repository/styleRepository"
)

func Add(style string) error {
	err := stylerepository.Add(style)

	return err
}

func GetAll() (*[]models.StyleDto, error) {
	styleList, err := stylerepository.GetAll()

	if err != nil {
		return nil, err
	}

	dtoList := []models.StyleDto{}

	for _, entity := range *styleList {
		dtoList = append(dtoList, *entity.ToStyleDTO())
	}

	return &dtoList, nil
}
