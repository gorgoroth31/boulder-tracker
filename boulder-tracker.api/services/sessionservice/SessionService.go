package sessionservice

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	SessionState "github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/enums"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/models"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/repository/sessionrepository"
)

func GetOrCreateInProgressSessionForUser(userId uuid.UUID) (*models.Session, error) {
	doesSessionExist, err := sessionrepository.ExistsLiveOrInProgressSessionForUser(userId)

	if err != nil {
		log.Fatal(err)
	}

	if doesSessionExist {
		return sessionrepository.GetLiveOrInProgressSessionForUser(userId)
	}

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	newSession := models.Session{SessionState: SessionState.InProgress}

	AddSession(&newSession)

	if err != nil {
		return nil, err
	}

	return sessionrepository.GetLiveOrInProgressSessionForUser(userId)
}

func AddSession(session *models.Session) error {
	err := sessionrepository.Add(session)
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
