package usercontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gofrs/uuid/v5"
	guid "github.com/google/uuid"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/models"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/services/userservice"
	"github.com/gorilla/mux"
)

func AddUserForPrincipal(w http.ResponseWriter, r *http.Request) {
	var userDto models.UserDto

	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if userDto.UserName == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/text")
		w.Write([]byte("Der Nutzername darf nicht leer sein."))
		return
	}

	fmt.Println("Principal: " + r.Header.Get("principal"))

	token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

	userDto.Principal = token.RegisteredClaims.Subject

	err = userservice.AddUser(&userDto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/text")
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(201)
}

func ExistsUserWithPrincipal(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

	exists, err := userservice.ExistsUserWithPrincipal(token.RegisteredClaims.Subject)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(exists)
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

func GetByPrincipalForLogin(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

	user, err := userservice.GetByPrincipal(token.RegisteredClaims.Subject)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		return
	}

	encoder := json.NewEncoder(w)

	encoder.Encode(user)
}
