package usercontroller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gofrs/uuid/v5"
	guid "github.com/google/uuid"
	"github.com/gorgoroth31/boulder-tracker/models"
	"github.com/gorgoroth31/boulder-tracker/services/userservice"
	"github.com/gorilla/mux"
)

func Add(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var userDto models.UserDto

	err := decoder.Decode(&userDto)
	if err != nil {
		panic(err)
	}
	userservice.AddUser(&userDto)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		log.Fatal("No id in the path")
	}

	userservice.Delete(guid.UUID(uuid.FromStringOrNil(id)))
}

func GetByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email, ok := vars["email"]

	if !ok {
		log.Fatal("No email in the path")
	}

	user, err := userservice.GetByEmail(email)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	encoder := json.NewEncoder(w)

	encoder.Encode(user)
}
