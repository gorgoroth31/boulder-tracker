package boulderservice

import (
	"log"

	"github.com/gorgoroth31/boulder-tracker/models"
	"github.com/gorgoroth31/boulder-tracker/repository/boulderrepository"
)

func Add(boulder *models.Boulder) error {
	err := boulderrepository.Add(boulder)

	if err != nil {
		log.Fatal(err)
	}
	return err
}
