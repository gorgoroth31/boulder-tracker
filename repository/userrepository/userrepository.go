package userrepository

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/gorgoroth31/boulder-tracker/db"
	"github.com/gorgoroth31/boulder-tracker/models"
)

func Add(user *models.User) error {
	// get db connection
	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	userId := uuid.New()

	stmt, err := database.Prepare("INSERT INTO user (Id, UserName, Email, IsDeleted) VALUES (?, ?, ?, ?);")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(userId, user.UserName, user.Email, false)

	if err != nil {
		return err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return err
	}

	return nil
}

func Delete(userId uuid.UUID) error {
	// get db connection
	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	stmt, err := database.Prepare("DELETE FROM user WHERE Id = ?;")

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

func GetByEmail(userEmail string) (*models.User, error) {
	database, err := db.CreateDatabase()

	if err != nil {
		fmt.Println("database connection failed")
	}
	defer database.Close()

	stmt, err := database.Prepare("SELECT Id, UserName, Email, IsDeleted FROM session where Email = ?")

	defer stmt.Close()

	var user *models.User

	err = stmt.QueryRow(userEmail).Scan(&user.Id, &user.UserName, &user.Email, &user.IsDeleted)
	if err != nil {
		return nil, err
	}

	return user, nil
}
