package styleRepository

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/db"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/models"
)

func Add(style string) error {
	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	stmt, err := database.Prepare("INSERT INTO style (Id, alias) VALUES (?, ?);")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(uuid.New(), style)

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

func GetAll() (*[]models.Style, error) {
	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	rows, err := database.Query("SELECT Id, alias FROM style;")

	if err != nil {
		return nil, err
	}

	var styleList []models.Style

	for rows.Next() {
		var diff models.Style
		if err := rows.Scan(&diff.Id, &diff.Alias); err != nil {
			return &styleList, err
		}
		styleList = append(styleList, diff)
	}

	if err = rows.Err(); err != nil {
		return &styleList, err
	}

	return &styleList, nil
}
