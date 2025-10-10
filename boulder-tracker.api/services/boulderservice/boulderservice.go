package boulderservice

import (
	"log"

	"github.com/google/uuid"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/models"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/repository/boulderrepository"
)

func Add(boulder *models.Boulder) error {
	err := boulderrepository.Add(boulder)

	if err != nil {
		log.Fatal(err)
	}
	return err
}

func GetBouldersForSessionId(sessionId uuid.UUID) (*[]models.Boulder, error) {
	return boulderrepository.GetBouldersForSessionId(sessionId)
}

func AddOrUpdate(boulder *models.Boulder) error {
	return boulderrepository.AddOrUpdate(boulder)
}

func DeleteById(boulderId uuid.UUID) error {
	return boulderrepository.DeleteById(boulderId)
}
