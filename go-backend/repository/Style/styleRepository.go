package style

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorgoroth31/boulder-tracker/go-backend/repository/Dbcontext"
)

func Add(style string) error {
	stmt, err := Dbcontext.DB.Prepare("INSERT INTO style (alias) VALUES (@alias);")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		sql.Named("alias", style),
	)

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
