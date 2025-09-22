package sessionservice

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/models"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/repository/sessionrepository"
)

func AddSession(session *models.SessionDto) error {
	sessionEntity := session.ToSessionEntity()
	err := sessionrepository.Add(sessionEntity)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func Delete(sessionId uuid.UUID) error {
	err := sessionrepository.Delete(sessionId)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetAllSessionsSimple() *[]models.SessionDto {
	sessions, err := sessionrepository.GetAllSessionsSimple()

	if err != nil {
		log.Fatal(err)
	}

	sessionsDTO := []models.SessionDto{}

	for _, session := range *sessions {
		sessionsDTO = append(sessionsDTO, *session.ToSessionDTO())
	}

	return &sessionsDTO
}
