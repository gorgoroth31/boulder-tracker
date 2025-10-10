package bouldercontroller

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/services/boulderservice"
	"github.com/gorilla/mux"
)

func DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		log.Fatal("No id in the path")
	}

	parsedUuid, err := uuid.ParseBytes([]byte(id))

	if err != nil {
		w.WriteHeader(400)
		return
	}

	// TODO restrict everyone except owner from deleting boulderroute

	err = boulderservice.DeleteById(parsedUuid)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(204)
}
