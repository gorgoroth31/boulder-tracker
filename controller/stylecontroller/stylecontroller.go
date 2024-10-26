package stylecontroller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorgoroth31/boulder-tracker/services/styleservice"
	"github.com/gorilla/mux"
)

func AddStyle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	alias, ok := vars["alias"]

	if !ok {
		log.Fatal("No alias in the path")
	}
	fmt.Println(alias)
	return
	err := styleservice.Add("newStyle")
	if err != nil {
		log.Fatal(err)
	}
}
