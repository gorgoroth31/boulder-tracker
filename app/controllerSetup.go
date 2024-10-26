package app

import (
	"github.com/gorgoroth31/boulder-tracker/controller/difficultycontroller"
	"github.com/gorgoroth31/boulder-tracker/controller/sessioncontroller"
	"github.com/gorgoroth31/boulder-tracker/controller/stylecontroller"
	"github.com/gorilla/mux"
)

func SetupController(router *mux.Router) {
	setupSessionController(router)
	setupStyleController(router)
	setupDifficultyController(router)
}

func setupSessionController(router *mux.Router) {
	router.Methods("POST").Path("/session").HandlerFunc(sessioncontroller.Add)
}

func setupStyleController(router *mux.Router) {
	router.Methods("POST").Path("/style/{alias}").HandlerFunc(stylecontroller.Add)
}

func setupDifficultyController(router *mux.Router) {
	router.Methods("POST").Path("/difficulty").HandlerFunc(difficultycontroller.Add)
}
