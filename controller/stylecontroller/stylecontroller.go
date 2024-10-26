package stylecontroller

import (
	"log"
	"net/http"

	styleservice "github.com/gorgoroth31/boulder-tracker/services/StyleService"
	"github.com/gorilla/mux"
)

func AddStyle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	alias, ok := vars["alias"]

	if !ok {
		log.Fatal("No alias in the path")
	}

	err := styleservice.Add(alias)
	if err != nil {
		log.Fatal(err)
	}
}
