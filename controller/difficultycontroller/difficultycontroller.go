package difficultycontroller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorgoroth31/boulder-tracker/models/dto"
)

func Add(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var t dto.DifficultyDto

	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Println(t)
	// vars := mux.Vars(r)
	// alias, ok := vars["alias"]

	// if !ok {
	// 	log.Fatal("No alias in the path")
	// }

	// err := difficultyservice.Add(alias)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
