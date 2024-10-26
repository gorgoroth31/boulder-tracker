package difficultycontroller

import (
	"encoding/json"
	"net/http"

	"github.com/gorgoroth31/boulder-tracker/models/dto"
	"github.com/gorgoroth31/boulder-tracker/services/difficultyservice"
)

func Add(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var t dto.DifficultyDto

	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	difficultyservice.Add(*t.ToDifficultyEntity())
}
