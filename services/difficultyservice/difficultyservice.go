package difficultyservice

import (
	"github.com/gorgoroth31/boulder-tracker/mapper"
	"github.com/gorgoroth31/boulder-tracker/models"
	"github.com/gorgoroth31/boulder-tracker/models/dto"
	"github.com/gorgoroth31/boulder-tracker/repository/difficultyrepository"
)

func Add(entity models.Difficulty) error {
	err := difficultyrepository.Add(entity)
	return err
}

func GetAll() (*[]dto.DifficultyDto, error) {
	difficultyList, err := difficultyrepository.GetAll()

	if err != nil {
		return nil, err
	}

	dtoList := []dto.DifficultyDto{}

	for _, entity := range *difficultyList {
		dtoList = append(dtoList, *mapper.ToDifficultyDto(&entity))
	}

	return &dtoList, nil
}
