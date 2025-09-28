package userrepository

import (
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/db"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/models"
)

func Add(user *models.User) error {
	if ExistsUsername(user.UserName) {
		err := errors.New("Dieser Name existiert bereits. Bitte wähle einen anderen.")
		return err
	}

	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	userId := uuid.New()

	stmt, err := database.Prepare("INSERT INTO users (Id, UserName, Principal, IsDeleted) VALUES (?, ?, ?, ?);")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(userId, user.UserName, user.Principal, false)

	if err != nil {
		return err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return err
	}

	return nil
}

func ExistsUserWithPrincipal(principal string) (bool, error) {
	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	stmt, err := database.Prepare("SELECT COUNT(*) FROM users WHERE Principal = ?")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var count int

	err = stmt.QueryRow(principal).Scan(&count)
	if err != nil {
		return count == 1, err
	}

	return count == 1, nil
}

func ExistsUsername(name string) bool {
	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	stmt, err := database.Prepare("SELECT COUNT(*) FROM users WHERE UserName = ?")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var count int

	err = stmt.QueryRow(name).Scan(&count)
	if err != nil {
		return count == 1
	}

	return count == 1
}

func Delete(userId uuid.UUID) error {
	// get db connection
	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	stmt, err := database.Prepare("DELETE FROM users WHERE Id = ?;")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(userId)

	if err != nil {
		return err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return err
	}

	return nil
}

func GetByPrincipal(principal string) (*models.User, error) {
	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	stmt, err := database.Prepare("SELECT Id, UserName, Principal, IsDeleted FROM users where Principal = ?")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var user models.User

	err = stmt.QueryRow(principal).Scan(&user.Id, &user.UserName, &user.Principal, &user.IsDeleted)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
