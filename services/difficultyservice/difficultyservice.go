package difficultyservice

import (
	"github.com/gorgoroth31/boulder-tracker/repository/difficultyrepository"
)

func Add(alias string) error {
	err := difficultyrepository.Add(alias)
	return err
}
