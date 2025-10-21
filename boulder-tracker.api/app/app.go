package app

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/app/middleware"
	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	Database *sql.DB
}

func (app *App) SetupRouter() {
	mount(app.Router, "/api/public", apiPublicRouter())
	mount(app.Router, "/api", apiPrivateRouter())
}

func apiPublicRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/health").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("health got pinged")
		w.Write([]byte("api is alive and well"))
	})
	return router
}

func apiPrivateRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	setupMiddleware(router)
	SetupController(router)
	return router
}

func setupMiddleware(router *mux.Router) {
	router.Use(middleware.SetupLogging)
	router.Use(middleware.AuthMiddleware)
}

func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}
