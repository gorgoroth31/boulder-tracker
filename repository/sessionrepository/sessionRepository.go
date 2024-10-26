package sessionrepository

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/gorgoroth31/boulder-tracker/db"
	"github.com/gorgoroth31/boulder-tracker/models"
	"github.com/gorgoroth31/boulder-tracker/services/boulderservice"
)

func Add(session *models.Session) error {
	// get db connection
	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	sessionId := uuid.New()

	stmt, err := database.Prepare("INSERT INTO session (Id, StartTime, EndTime, BoulderedSolo) VALUES (?, ?, ?, ?);")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(sessionId, session.VisitTime.From(), session.VisitTime.To(), session.BoulderedSolo)

	if err != nil {
		return err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return err
	}

	for _, boulder := range session.RoutesSolved {
		boulder.SessionId = sessionId
		err := boulderservice.Add(&boulder)
		if err != nil {
			return err
		}
	}

	return nil
}

func Delete(sessionId uuid.UUID) error {
	// get db connection
	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	stmt, err := database.Prepare("DELETE FROM session WHERE Id = ?;")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(sessionId)

	if err != nil {
		return err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return err
	}

	return nil
}
