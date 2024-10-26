package sessionservice

import (
	"log"

	"github.com/gorgoroth31/boulder-tracker/models/dto"
	"github.com/gorgoroth31/boulder-tracker/repository/sessionrepository"
)

func AddSession(session *dto.SessionDto) error {
	sessionEntity := session.ToSessionEntity()
	err := sessionrepository.Add(sessionEntity)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
