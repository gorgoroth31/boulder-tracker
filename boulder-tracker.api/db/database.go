package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func CreateDatabase() (*sql.DB, error) {

	User := os.Getenv("DBUSER")
	Passwd := os.Getenv("DBPASS")
	Net := "tcp"
	Addr := os.Getenv("DBURL")
	DBName := "bouldertracker"
	dsn := (User + ":" + Passwd + "@" + Net + "(" + Addr + ")/" + DBName + "?charset=utf8mb4&parseTime=True&loc=Local")

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(3)
	db.Exec("SET SQL_MODE='ALLOW_INVALID_DATES';")
	return db, nil
}
