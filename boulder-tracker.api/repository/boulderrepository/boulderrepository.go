package boulderrepository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/db"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/models"
)

func Add(boulder *models.Boulder) error {

	// get db connection
	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	stmt, err := database.Prepare("INSERT INTO boulders (Id, attempts, sessionsTried, exhausting, SessionId) VALUES (?, ?, ?, ?, ?);")

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

	if err != nil {
		return err
	}

	boulder.Id = boulderId

	err = addDifficulties(boulder, database)

	if err != nil {
		return err
	}

	return nil
}

func GetBouldersForSessionId(sessionId uuid.UUID) (*[]models.Boulder, error) {
	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	rows, err := database.Query("SELECT Id, ScrewedDifficulty, FeltLikeDifficulty, Attempts, SessionsTried, Exhausting, Style, Like, SessionId FROM boulders WHERE SessionId = ?;", sessionId)

	if err != nil {
		return nil, err
	}

	var routes []models.Boulder

	for rows.Next() {
		var diff models.Boulder
		if err := rows.Scan(&diff.Id, &diff.ScrewedDifficulty, &diff.FeltLikeDifficulty, &diff.Attempts, &diff.SessionsTried, &diff.Exhausting, &diff.Style, &diff.Like, &diff.SessionId); err != nil {
			return &routes, err
		}
		routes = append(routes, diff)
	}

	if err = rows.Err(); err != nil {
		return &routes, err
	}

	return &routes, nil
}

func Update(boulder *models.Boulder) error {
	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	stmt, err := database.Prepare("UPDATE boulders SET Attempts = ?, SessionsTried = ?, Exhausting = ? WHERE Id = ?;")

	if err != nil {
		return err
	}

	result, err := stmt.Exec(boulder.Attempts, boulder.SessionsTried, boulder.Exhausting, boulder.Id)

	if err != nil {
		return err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return err
	}

	err = updateStylesForBoulder(boulder)

	if err != nil {
		return err
	}

	err = updateDifficultyForBoulder(boulder)

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

func updateDifficultyForBoulder(boulder *models.Boulder) error {
	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	stmt, err := database.Prepare("UPDATE boulder_feltlike_difficulty SET DifficultyId = ? WHERE BoulderId = ?;")

	if err != nil {
		return err
	}

	result, err := stmt.Exec(boulder.FeltLikeDifficulty.Id, boulder.Id)

	if err != nil {
		return err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return err
	}

	stmt2, err := database.Prepare("UPDATE boulder_actual_difficulty Set DifficultyId = ? Where BoulderId = ?;")

	if err != nil {
		return err
	}

	result2, err := stmt2.Exec(boulder.ScrewedDifficulty.Id, boulder.Id)

	if err != nil {
		return err
	}

	_, err = result2.RowsAffected()

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

func addDifficulties(boulder *models.Boulder, db *sql.DB) error {

	err := addActualDifficulty(boulder, db)

	if err != nil {
		return err
	}

	err = addFeltLikeDifficulty(boulder, db)

	if err != nil {
		return err
	}

	return nil
}

func addActualDifficulty(boulder *models.Boulder, db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO boulder_actual_difficulty (boulderId, difficultyId) VALUES (?, ?);")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(boulder.Id, boulder.ScrewedDifficulty.Id)

	if err != nil {
		return err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return err
	}

	return nil
}

func addFeltLikeDifficulty(boulder *models.Boulder, db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO boulder_feltlike_difficulty (boulderId, difficultyId) VALUES (?, ?);")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(boulder.Id, boulder.FeltLikeDifficulty.Id)

	if err != nil {
		return err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return err
	}

	return nil
}
