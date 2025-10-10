package sessionservice

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	SessionState "github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/enums"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/models"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/repository/sessionrepository"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/services/userservice"
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

	newSession := models.Session{SessionState: SessionState.InProgress, UserId: userId, IsDeleted: false, StartTime: time.Now().Add(-300), EndTime: time.Now()}

	AddSession(&newSession)

	if err != nil {
		return nil, err
	}

	return sessionrepository.GetLiveOrInProgressSessionForUser(userId)
}

func GetLatestForUser(principalId string, count int) (*[]models.Session, error) {
	user, err := userservice.GetByPrincipal(principalId)

	if err != nil {
		return nil, err
	}

	return sessionrepository.GetLatestForUser(user.Id, count)
}

func Update(session *models.Session) (*models.Session, error) {
	err := sessionrepository.Update(session)

	if err != nil {
		return nil, err
	}
	return sessionrepository.GetSessionById(session.Id)
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

func SubmitCurrentSessionForUser(principalId string) error {
	user, err := userservice.GetByPrincipal(principalId)

	if err != nil {
		return err
	}

	doesSessionExist, err := sessionrepository.ExistsLiveOrInProgressSessionForUser(user.Id)

	if err != nil {
		return err
	}

	if !doesSessionExist {
		return nil
	}

	return sessionrepository.SubmitCurrentSessionForUserId(user.Id)
}
