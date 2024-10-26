package sessioncontroller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorgoroth31/boulder-tracker/models/dto"
)

func Add(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t dto.SessionDto
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Println(t)
}
