package difficultyrepository

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/gorgoroth31/boulder-tracker/db"
)

func Add(alias string) error {

	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	stmt, err := database.Prepare("INSERT INTO difficulty (Id, alias) VALUES (?, ?);")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(uuid.New(), alias)

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
