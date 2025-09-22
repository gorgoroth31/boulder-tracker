package usercontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gofrs/uuid/v5"
	guid "github.com/google/uuid"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/models"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/services/userservice"
	"github.com/gorilla/mux"
)

func Add(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var userDto models.UserDto

	err := decoder.Decode(&userDto)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	err = userservice.AddUser(&userDto)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(409)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		log.Fatal("No id in the path")
	}

	err := userservice.Delete(guid.UUID(uuid.FromStringOrNil(id)))

	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(201)
}

func GetByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email, ok := vars["email"]

	if !ok {
		fmt.Println("No email in the path")
		w.WriteHeader(403)
		return
	}

	user, err := userservice.GetByEmail(email)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		return
	}

	encoder := json.NewEncoder(w)

	encoder.Encode(user)
}
