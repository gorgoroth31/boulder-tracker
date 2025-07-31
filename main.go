package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"slices"

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
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(app.Router)

	fmt.Println("BoulderTracker.API is now listening at " + ipWithPort)

	log.Fatal(http.ListenAndServe(ipWithPort, handler))
}
