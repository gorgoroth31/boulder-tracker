package app

import (
	"net/http"

	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/controller/bouldercontroller"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/controller/difficultycontroller"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/controller/sessioncontroller"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/controller/stylecontroller"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/controller/usercontroller"
	"github.com/gorilla/mux"
)

func SetupController(router *mux.Router) {
	setupHealthController(router)
	setupUserController(router)
	setupSessionController(router)
	setupStyleController(router)
	setupDifficultyController(router)
	setupBoulderController(router)
}

func setupUserController(router *mux.Router) {
	router.HandleFunc("/user/login", usercontroller.AddUserForPrincipal).Methods("POST")
	router.HandleFunc("/user/exists", usercontroller.ExistsUserWithPrincipal).Methods("GET")
	router.HandleFunc("/user/login", usercontroller.GetByPrincipalForLogin).Methods("GET")
}

func setupSessionController(router *mux.Router) {
	router.HandleFunc("/session/currentLiveOrInProgress", sessioncontroller.GetCurrentSession).Methods("GET")
	router.HandleFunc("/session/getLatest", sessioncontroller.GetLatest).Methods("GET")
	router.HandleFunc("/session", sessioncontroller.UpdateSession).Methods("PUT")
	router.HandleFunc("/session/submitCurrent", sessioncontroller.SubmitCurrentSession).Methods("PUT")
}

func setupStyleController(router *mux.Router) {
	router.HandleFunc("/style/{alias}", stylecontroller.Add).Methods("POST")
	router.HandleFunc("/style", stylecontroller.GetAll).Methods("GET")
}

func setupDifficultyController(router *mux.Router) {
	router.HandleFunc("/difficulty", difficultycontroller.Add).Methods("POST")
	router.HandleFunc("/difficulty", difficultycontroller.GetAll).Methods("GET")
}

func setupBoulderController(router *mux.Router) {
	router.HandleFunc("/boulder/{id}", bouldercontroller.DeleteById).Methods("DELETE")
}

func setupHealthController(router *mux.Router) {
	router.HandleFunc("/health", health).Methods("GET")
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("api is alive and well"))
}
