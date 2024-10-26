package main

import (
	"log"
	"net/http"

	styleRepository "github.com/gorgoroth31/boulder-tracker/repository/StyleRepository"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	setupRouter(router)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))
}

func setupRouter(router *mux.Router) {
	router.Methods("POST").Path("/style").HandlerFunc(addStyle)
}

func addStyle(w http.ResponseWriter, r *http.Request) {
	err := styleRepository.Add("newStyle")
	if err != nil {
		log.Fatal(err)
	}
}
