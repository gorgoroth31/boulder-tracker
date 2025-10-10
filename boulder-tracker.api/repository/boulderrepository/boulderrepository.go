package boulderrepository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/db"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/models"
	styleservice "github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/services/StyleService"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/services/difficultyservice"
)

func Add(boulder *models.Boulder) error {

	// get db connection
	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	stmt, err := database.Prepare("INSERT INTO boulders (Id, Attempts, SessionsTried, Exhausting, SessionId) VALUES (?, ?, ?, ?, ?);")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	boulderId := uuid.New()

	result, err := stmt.Exec(boulderId, boulder.Attempts, boulder.SessionsTried, boulder.Exhausting, boulder.SessionId)

	if err != nil {
		return err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return err
	}

	err = addStyles(boulderId, boulder.Style, database)

	return err
}

func GetBouldersForSessionId(sessionId uuid.UUID) (*[]models.Boulder, error) {
	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	rows, err := database.Query("SELECT Id, ScrewedDifficultyId, FeltLikeDifficultyId, Attempts, SessionsTried, Exhausting, Likes, SessionId FROM boulders WHERE SessionId = ?;", sessionId)

	if err != nil {
		return nil, err
	}

	var routes []models.Boulder

	for rows.Next() {
		var diff models.Boulder
		if err := rows.Scan(&diff.Id, &diff.ScrewedDifficultyId, &diff.FeltLikeDifficultyId, &diff.Attempts, &diff.SessionsTried, &diff.Exhausting, &diff.Like, &diff.SessionId); err != nil {
			return &routes, err
		}

		feltLikeDifficulty, err := difficultyservice.GetById(diff.FeltLikeDifficultyId)

		if err != nil {
			return nil, err
		}

		diff.FeltLikeDifficulty = *feltLikeDifficulty

		screwedDifficulty, err := difficultyservice.GetById(diff.ScrewedDifficultyId)

		if err != nil {
			return nil, err
		}

		diff.ScrewedDifficulty = *screwedDifficulty

		styles, err := styleservice.GetAllByBoulderId(diff.Id)

		if err != nil {
			return nil, err
		}

		diff.Style = append(diff.Style, *styles...)

		routes = append(routes, diff)
	}

	if err = rows.Err(); err != nil {
		return &routes, err
	}

	return &routes, nil
}

func AddOrUpdate(boulder *models.Boulder) error {
	if boulder.Id == uuid.Nil {
		return addBoulder(boulder)
	}

	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	stmt, err := database.Prepare("UPDATE boulders SET Attempts = ?, SessionsTried = ?, Exhausting = ?, FeltLikeDifficultyId = ?, ScrewedDifficultyId = ? WHERE Id = ?;")

	if err != nil {
		return err
	}

	result, err := stmt.Exec(boulder.Attempts, boulder.SessionsTried, boulder.Exhausting, boulder.FeltLikeDifficulty.Id, boulder.ScrewedDifficulty.Id, boulder.Id)

	if err != nil {
		return err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return err
	}

	err = updateStylesForBoulder(boulder)

	return err
}

func DeleteById(boulderId uuid.UUID) error {
	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
		return err
	}
	defer database.Close()

	stmt, err := database.Prepare("DELETE FROM boulders WHERE Id = ?;")

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(boulderId)

	return err
}

func addBoulder(boulder *models.Boulder) error {
	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	stmt, err := database.Prepare("INSERT INTO boulders (Id, Attempts, SessionsTried, Exhausting, Likes, FeltLikeDifficultyId, ScrewedDifficultyId, SessionId) VALUES (?, ?, ?, ?, ?, ?, ?, ?);")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	boulder.Id = uuid.New()

	result, err := stmt.Exec(boulder.Id, boulder.Attempts, boulder.SessionsTried, boulder.Exhausting, boulder.Like, boulder.FeltLikeDifficulty.Id, boulder.ScrewedDifficulty.Id, boulder.SessionId)

	if err != nil {
		return err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return err
	}

	err = updateStylesForBoulder(boulder)

	return err
}

func updateStylesForBoulder(boulder *models.Boulder) error {
	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	stmt, err := database.Prepare("DELETE from boulder_styles WHERE BoulderId = ?;")

	if err != nil {
		return err
	}

	result, err := stmt.Exec(boulder.Id)

	if err != nil {
		return err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return err
	}

	for _, style := range boulder.Style {
		err = addStyleToBoulderId(boulder.Id, style.Id)

		if err != nil {
			return err
		}
	}

	return nil
}

func addStyleToBoulderId(boulderId uuid.UUID, styleId uuid.UUID) error {
	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	stmt, err := database.Prepare("INSERT INTO boulder_styles (BoulderId, StyleId) VALUES (?, ?);")

	if err != nil {
		return err
	}

	result, err := stmt.Exec(boulderId, styleId)

	if err != nil {
		return err
	}

	_, err = result.RowsAffected()

	return err
}

func addStyles(boulderId uuid.UUID, styles []models.Style, db *sql.DB) error {

	for _, style := range styles {
		stmt, err := db.Prepare("INSERT INTO boulder_styles (boulderId, styleId) VALUES (?, ?);")
		if err != nil {
			log.Fatal(err)
		}

		result, err := stmt.Exec(boulderId, style.Id)

		if err != nil {
			return err
		}

		_, err = result.RowsAffected()
		if err != nil {
			return err
		}
		stmt.Close()
	}

	return nil
}
