package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	setupRouter(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func setupRouter(router *mux.Router) {
	router.Methods("POST").Path("/endpoint").HandlerFunc(postFunction)
}

func postFunction(w http.ResponseWriter, r *http.Request) {
	log.Print("You called that")
}
