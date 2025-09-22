package app

import (
	"fmt"
	"net/http"

	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/controller/difficultycontroller"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/controller/sessioncontroller"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/controller/stylecontroller"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/controller/usercontroller"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/db"
	"github.com/gorilla/mux"
)

func SetupController(router *mux.Router) {
	setupUserController(router)
	setupSessionController(router)
	setupStyleController(router)
	setupDifficultyController(router)
	initDifficultyAndStyleController(router)
	setupHealthController(router)
}

func setupUserController(router *mux.Router) {
	router.Methods("POST").Path("/user").HandlerFunc(usercontroller.Add)
	router.Methods("DELETE").Path("/user/{id}").HandlerFunc(usercontroller.Delete)
	router.Methods("GET").Path("/user/byEmail/{email}").HandlerFunc(usercontroller.GetByEmail)
}

func setupSessionController(router *mux.Router) {
	router.Methods("POST").Path("/session").HandlerFunc(sessioncontroller.Add)
	router.Methods("DELETE").Path("/session/{id}").HandlerFunc(sessioncontroller.Delete)
	router.Methods("GET").Path("/session").HandlerFunc(sessioncontroller.GetAllSessionsSimple)
}

func setupStyleController(router *mux.Router) {
	router.Methods("POST").Path("/style/{alias}").HandlerFunc(stylecontroller.Add)
	router.Methods("GET").Path("/style").HandlerFunc(stylecontroller.GetAll)
}

func setupDifficultyController(router *mux.Router) {
	router.Methods("POST").Path("/difficulty").HandlerFunc(difficultycontroller.Add)
	router.Methods("GET").Path("/difficulty").HandlerFunc(difficultycontroller.GetAll)
}

func initDifficultyAndStyleController(router *mux.Router) {
	router.Methods("POST").Path("/init").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := db.CreateDatabase()
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = db.Exec(
			`
INSERT INTO difficulty (Id, alias, relativeLevel)
VALUES
('e87067fd-71ef-4b5a-8e4d-2d328be9a894', 'gelb', '0'),
('fd0f9dcb-72b4-4b83-ad9f-13519fa0f545', 'grün', '1'),
('3cbdaefc-b51d-4a9d-a6d6-bf47b536bcea', 'orange', '2'),
('139371d9-e707-4bfe-b3c9-16c0a5535d1d', 'weiß', '3'),
('26e2e5b9-b1fd-40fa-a1c8-c44feadda09f', 'blau', '4'),
('2797f7fb-0655-4888-8525-5b1c4abb88bc', 'rot', '5'),
('dc34cea4-4d7f-4eb0-bbe5-bbccd7e3c8c1', 'schwarz', '6');
`,
		)

		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = db.Exec(
			`
INSERT INTO style (Id, alias) 
VALUES 
('e9de5a39-4ca0-4f15-b105-99c87b529a1c', 'slab'),
('5174a6e7-0f7d-4fff-a7ab-88b9c08f2fb7', 'überhang'),
('c0a26d92-c1ee-4be6-a778-5ef6d13e42dd', 'slopey'),
('382dc5a4-45f2-49d0-8457-e886dd223240', 'lang'),
('25da5be5-5857-4e30-9372-a95c502a96e7', 'kurz'),
('c704d613-e173-4202-bad1-7ecf739b3cc7', 'statisch'),
('9667d906-8c07-4209-87f0-f3f9096a0f17', 'dynamisch'),
('21d6312a-1bce-4541-92de-5e118b829789', 'stretchy'),
('1aa3ddba-7b1c-45a0-9a67-554e8ef0178f', 'kraftig'),
('79031454-5f7a-4139-b927-31da474b203d', 'einfach hoch');


`,
		)

		if err != nil {
			fmt.Println(err)
			return
		}
	})
}

func setupHealthController(router *mux.Router) {
	router.Methods("GET").Path("/health").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("health got pinged")
		w.Write([]byte("api is alive and well"))
	})
}
