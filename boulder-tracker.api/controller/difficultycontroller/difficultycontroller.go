package difficultycontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/models"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/services/difficultyservice"
)

func Add(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var t models.DifficultyDto

	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	difficultyservice.Add(*t.ToDifficultyEntity())
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	list, err := difficultyservice.GetAll()

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	encoder := json.NewEncoder(w)

	encoder.Encode(list)
}
