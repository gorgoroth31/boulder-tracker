package app

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gorgoroth31/boulder-tracker/controller/stylecontroller"
	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	Database *sql.DB
}

func (app *App) SetupRouter() {
	mount(app.Router, "/api", apiRouter())
}

func apiRouter() *mux.Router {
	router := mux.NewRouter()
	setupMiddleware(router)
	setupSessionController(router)
	setupStyleController(router)
	return router
}

func setupMiddleware(router *mux.Router) {
	router.Use(LoggingMiddleware)
}

func setupSessionController(router *mux.Router) {
	router.Methods("POST").Path("/session")
}

func setupStyleController(router *mux.Router) {
	router.Methods("POST").Path("/style/{alias}").HandlerFunc(stylecontroller.AddStyle)
}

func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}
