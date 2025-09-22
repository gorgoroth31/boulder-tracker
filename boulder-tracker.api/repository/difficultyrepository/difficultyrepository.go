package difficultyrepository

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/db"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/models"
)

func Add(entity models.Difficulty) error {

	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	stmt, err := database.Prepare("INSERT INTO difficulties (Id, alias, relativeLevel) VALUES (?, ?, ?);")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(uuid.New(), entity.Alias, entity.RelativeLevel)

	if err != nil {
		return err
	}

	rowInserted, err := result.RowsAffected()

	fmt.Println(rowInserted)

	if err != nil {
		return err
	}

	return nil
}

func GetAll() (*[]models.Difficulty, error) {
	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	rows, err := database.Query("SELECT * FROM difficulties;")

	if err != nil {
		return nil, err
	}

	var difficultyList []models.Difficulty

	for rows.Next() {
		var diff models.Difficulty
		if err := rows.Scan(&diff.Id, &diff.Alias, &diff.RelativeLevel); err != nil {
			return &difficultyList, err
		}
		difficultyList = append(difficultyList, diff)
	}

	if err = rows.Err(); err != nil {
		return &difficultyList, err
	}

	return &difficultyList, nil
}
