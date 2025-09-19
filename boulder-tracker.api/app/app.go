package app

import (
	"database/sql"
	"net/http"
	"strings"

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
	router := mux.NewRouter().StrictSlash(true)
	setupMiddleware(router)
	SetupController(router)
	return router
}

func setupMiddleware(router *mux.Router) {
	router.Use(LoggingMiddleware)
}

func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}
