package app

import (
	"database/sql"

	"github.com/gorgoroth31/boulder-tracker/controller/stylecontroller"
	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	Database *sql.DB
}

func (app *App) SetupRouter() {
	app.setupStyleController()
}

func (app *App) setupStyleController() {
	app.Router.Methods("POST").Path("/style").HandlerFunc(stylecontroller.AddStyle)
}
