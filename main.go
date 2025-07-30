package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorgoroth31/boulder-tracker/app"
	"github.com/gorgoroth31/boulder-tracker/db"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	database, err := db.CreateDatabase()
	if err != nil {
		log.Fatal("Database connection failed:")
	}

	app := &app.App{
		Router:   mux.NewRouter(),
		Database: database,
	}

	app.SetupRouter()

	// TODO: move to env variable
	ipAddress := "127.0.0.1"
	port := "8080"

	ipWithPort := ipAddress + ":" + port

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "DELETE", "GET"},
		AllowCredentials: true,
	})

	handler := c.Handler(app.Router)

	fmt.Println("BoulderTracker.API is now listening at " + ipWithPort)

	log.Fatal(http.ListenAndServe(ipWithPort, handler))
}
