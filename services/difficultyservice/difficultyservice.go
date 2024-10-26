package difficultyservice

import (
	"github.com/gorgoroth31/boulder-tracker/models"
	"github.com/gorgoroth31/boulder-tracker/repository/difficultyrepository"
)

func Add(entity models.Difficulty) error {
	err := difficultyrepository.Add(entity)
	return err
}
