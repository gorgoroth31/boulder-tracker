package sessioncontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/models"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/services/sessionservice"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/services/userservice"
)

func GetCurrentSession(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

	user, err := userservice.GetByPrincipal(token.RegisteredClaims.Subject)

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

func UpdateSession(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var sessionDto models.SessionDto

	err := decoder.Decode(&sessionDto)
	if err != nil {
		panic(err)
	}

	sessionEntity := sessionDto.ToSessionEntity()

	err = sessionservice.Update(sessionEntity)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}
