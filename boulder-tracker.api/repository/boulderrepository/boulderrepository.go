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

	stmt, err := database.Prepare("INSERT INTO boulder (Id, attempts, sessionsTried, exhausting, SessionId) VALUES (?, ?, ?, ?, ?);")

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

func addStyles(boulderId uuid.UUID, styles []models.Style, db *sql.DB) error {

	for _, style := range styles {
		stmt, err := db.Prepare("INSERT INTO boulder_style (boulderId, styleId) VALUES (?, ?);")
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
