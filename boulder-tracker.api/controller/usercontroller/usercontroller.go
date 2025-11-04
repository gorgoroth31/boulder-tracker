package usercontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/models"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/services/userservice"
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

	principal := r.Header.Get("principal")

	userDto.Principal = principal

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
	principal := r.Header.Get("principal")

	exists, err := userservice.ExistsUserWithPrincipal(principal)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(exists)
}

func GetByPrincipalForLogin(w http.ResponseWriter, r *http.Request) {
	principal := r.Header.Get("principal")

	user, err := userservice.GetByPrincipal(principal)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		return
	}

	encoder := json.NewEncoder(w)

	encoder.Encode(user)
}
