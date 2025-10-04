package sessionrepository

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/db"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/models"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/repository/boulderrepository"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/services/boulderservice"
)

func ExistsLiveOrInProgressSessionForUser(userId uuid.UUID) (bool, error) {
	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	stmt, err := database.Prepare("SELECT COUNT(*) FROM sessions WHERE UserId = ? AND SessionState IN (0,1);")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var count int

	err = stmt.QueryRow(userId).Scan(&count)
	if err != nil {
		return count == 1, err
	}

	return count == 1, nil
}

func GetLiveOrInProgressSessionForUser(userId uuid.UUID) (*models.Session, error) {
	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	stmt, err := database.Prepare("SELECT Id, StartTime, EndTime, BoulderedSolo, UserId, IsDeleted, SessionState FROM sessions WHERE UserId = ? AND SessionState IN (0, 1);")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var session models.Session

	err = stmt.QueryRow(userId).Scan(&session.Id, &session.StartTime, &session.EndTime, &session.BoulderedSolo, &session.UserId, &session.IsDeleted, &session.SessionState)
	if err != nil {
		return nil, err
	}

	// fill in rest of fields
	routes, err := boulderrepository.GetBouldersForSessionId(session.Id)

	if err != nil {
		return nil, err
	}

	session.RoutesSolved = *routes

	return &session, nil
}

func Add(session *models.Session) error {
	// get db connection
	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	sessionId := uuid.New()

	stmt, err := database.Prepare("INSERT INTO sessions (Id, StartTime, EndTime, BoulderedSolo, SessionState, UserId) VALUES (?, ?, ?, ?, ?, ?);")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(sessionId, session.StartTime, session.EndTime, session.BoulderedSolo, session.SessionState, session.UserId)

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

	stmt, err := database.Prepare("DELETE FROM sessions WHERE Id = ?;")

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

func GetAllSessionsSimple() (*[]models.Session, error) {
	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	rows, err := database.Query("SELECT Id, StartTime, EndTime, SessionState FROM sessions;")

	if err != nil {
		return nil, err
	}

	var sessions []models.Session

	for rows.Next() {
		var diff models.Session
		if err := rows.Scan(&diff.Id, &diff.StartTime, &diff.EndTime, &diff.SessionState); err != nil {
			return &sessions, err
		}
		sessions = append(sessions, diff)
	}

	if err = rows.Err(); err != nil {
		return &sessions, err
	}

	return &sessions, nil
}
