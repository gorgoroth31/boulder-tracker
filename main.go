package main

import (
	"log"
	"net/http"

	"github.com/gorgoroth31/boulder-tracker/app"
	"github.com/gorgoroth31/boulder-tracker/db"
	"github.com/gorilla/mux"
)

func main() {
	database, err := db.CreateDatabase()
	if err != nil {
		log.Fatal("Database connection failed:")
	}

	app := &app.App{
		Router:   mux.NewRouter().StrictSlash(true),
		Database: database,
	}

	app.SetupRouter()

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", app.Router))
}
