package difficultyservice

import (
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/models"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/repository/difficultyrepository"
)

func Add(entity models.Difficulty) error {
	err := difficultyrepository.Add(entity)
	return err
}

func GetAll() (*[]models.DifficultyDto, error) {
	difficultyList, err := difficultyrepository.GetAll()

	if err != nil {
		return nil, err
	}

	dtoList := []models.DifficultyDto{}

	for _, entity := range *difficultyList {
		dtoList = append(dtoList, *entity.ToDifficultyDto())
	}

	return &dtoList, nil
}
