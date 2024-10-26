package difficultycontroller

import (
	"fmt"
	"net/http"
)

func Add(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Body)

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
