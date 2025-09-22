package stylecontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	styleservice "github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/services/StyleService"
	"github.com/gorilla/mux"
)

func Add(w http.ResponseWriter, r *http.Request) {
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

func GetAll(w http.ResponseWriter, r *http.Request) {
	list, err := styleservice.GetAll()

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	encoder := json.NewEncoder(w)

	encoder.Encode(list)
}
