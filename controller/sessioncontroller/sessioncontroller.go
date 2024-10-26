package sessioncontroller

import (
	"encoding/json"
	"net/http"

	"github.com/gorgoroth31/boulder-tracker/models/dto"
	"github.com/gorgoroth31/boulder-tracker/services/sessionservice"
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
