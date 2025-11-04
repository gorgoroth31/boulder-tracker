package sessioncontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/models"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/services/sessionservice"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/services/userservice"
)

func GetCurrentSession(w http.ResponseWriter, r *http.Request) {
	principal := r.Header.Get("principal")

	user, err := userservice.GetByPrincipal(principal)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	session, err := sessionservice.GetOrCreateInProgressSessionForUser(user.Id)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}

	encoder := json.NewEncoder(w)

	encoder.Encode(session.ToSessionDTO())
}

func GetLatest(w http.ResponseWriter, r *http.Request) {
	principal := r.Header.Get("principal")

	sessions, err := sessionservice.GetLatestForUser(principal, 5)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}

	var sessionDtos []models.SessionDto

	for _, v := range *sessions {
		sessionDtos = append(sessionDtos, *v.ToSessionDTO())
	}

	encoder := json.NewEncoder(w)

	encoder.Encode(sessionDtos)
}

func UpdateSession(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var sessionDto models.SessionDto

	err := decoder.Decode(&sessionDto)
	if err != nil {
		panic(err)
	}

	// TODO restrict users, who are not the owner from updating

	sessionEntity := sessionDto.ToSessionEntity()

	session, err := sessionservice.Update(sessionEntity)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}

	encoder := json.NewEncoder(w)

	encoder.Encode(session.ToSessionDTO())
}

func SubmitCurrentSession(w http.ResponseWriter, r *http.Request) {
	principal := r.Header.Get("principal")

	err := sessionservice.SubmitCurrentSessionForUser(principal)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}
