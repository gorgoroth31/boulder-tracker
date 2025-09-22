package sessioncontroller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gofrs/uuid/v5"
	guid "github.com/google/uuid"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/models"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/services/sessionservice"
	"github.com/gorilla/mux"
)

func Add(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var sessionDto models.SessionDto

	err := decoder.Decode(&sessionDto)
	if err != nil {
		panic(err)
	}

	if sessionDto.UserId == guid.UUID(uuid.Nil) {
		w.WriteHeader(400)
		return
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

func GetAllSessionsSimple(w http.ResponseWriter, r *http.Request) {
	sessions := sessionservice.GetAllSessionsSimple()

	encoder := json.NewEncoder(w)

	encoder.Encode(sessions)
}
