package sessioncontroller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gofrs/uuid/v5"
	guid "github.com/google/uuid"
	"github.com/gorgoroth31/boulder-tracker/models/dto"
	"github.com/gorgoroth31/boulder-tracker/services/sessionservice"
	"github.com/gorilla/mux"
)

func Add(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var sessionDto dto.SessionDto

	err := decoder.Decode(&sessionDto)
	if err != nil {
		panic(err)
	}
	sessionservice.AddSession(&sessionDto)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		log.Fatal("No id in the path")
	}

	sessionservice.Delete(guid.UUID(uuid.FromStringOrNil(id)))
}
