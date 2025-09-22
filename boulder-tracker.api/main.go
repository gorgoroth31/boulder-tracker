package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"slices"

	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/app"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/db"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error loading .env file")
	}

	database, err := db.CreateDatabase()
	if err != nil {
		log.Fatal("Database connection failed:")
	}

	app := &app.App{
		Router:   mux.NewRouter(),
		Database: database,
	}

	app.SetupRouter()

	ipAddress := "localhost"
	port := "8080"

	if len(os.Args) > 1 {
		argsWithoutProg := os.Args[1:]
		if slices.Contains(argsWithoutProg, "--host") {
			for i, v := range argsWithoutProg {
				if v == "--host" {
					ipAddress = argsWithoutProg[i+1]
					break
				}
			}
		}
	}

	ipWithPort := ipAddress + ":" + port

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "DELETE", "POST", "PUT"},
		AllowedHeaders:   []string{"Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(app.Router)

	fmt.Println("BoulderTracker.API is now listening at " + ipWithPort)

	log.Fatal(http.ListenAndServe(ipWithPort, handler))
}
