package main

import (
	"fmt"
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

	ipAddress := "127.0.0.1"
	port := "8080"

	ipWithPort := ipAddress + ":" + port

	fmt.Println("BoulderTracker.API is now listening at " + ipWithPort)

	log.Fatal(http.ListenAndServe(ipWithPort, app.Router))
}
