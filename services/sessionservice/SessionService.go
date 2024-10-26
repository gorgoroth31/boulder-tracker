package sessionservice

import (
	"fmt"
	"log"

	"github.com/google/uuid"
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

func Delete(sessionId uuid.UUID) error {
	err := sessionrepository.Delete(sessionId)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
