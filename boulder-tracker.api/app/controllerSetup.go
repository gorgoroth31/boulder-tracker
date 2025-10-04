package app

import (
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/app/middleware"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/controller/difficultycontroller"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/controller/sessioncontroller"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/controller/stylecontroller"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/controller/usercontroller"
	"github.com/gorilla/mux"
)

func SetupController(router *mux.Router) {
	setupHealthController(router)
	setupUserController(router)
	setupSessionController(router)
	setupStyleController(router)
	setupDifficultyController(router)
}

func setupUserController(router *mux.Router) {
	router.Handle("/user/login", middleware.EnsureValidToken()(baseHandlerFunc(usercontroller.AddUserForPrincipal, ""))).Methods("POST")
	router.Handle("/user/exists", middleware.EnsureValidToken()(baseHandlerFunc(usercontroller.ExistsUserWithPrincipal, ""))).Methods("GET")
	router.Handle("/user/{id}", middleware.EnsureValidToken()(baseHandlerFunc(usercontroller.Delete, ""))).Methods("DELETE")
	router.Handle("/user/login", middleware.EnsureValidToken()(baseHandlerFunc(usercontroller.GetByPrincipalForLogin, ""))).Methods("GET")
}

func setupSessionController(router *mux.Router) {
	router.Handle("/session/currentLiveOrInProgress", middleware.EnsureValidToken()(baseHandlerFunc(sessioncontroller.GetCurrentSession, ""))).Methods("GET")
	router.Handle("/session", middleware.EnsureValidToken()(baseHandlerFunc(sessioncontroller.UpdateSession, ""))).Methods("PUT")
}

func setupStyleController(router *mux.Router) {
	router.Handle("/style/{alias}", middleware.EnsureValidToken()(baseHandlerFunc(stylecontroller.Add, ""))).Methods("POST")
	router.Handle("/style", middleware.EnsureValidToken()(baseHandlerFunc(stylecontroller.GetAll, ""))).Methods("GET")
}

func setupDifficultyController(router *mux.Router) {
	router.Handle("/difficulty", middleware.EnsureValidToken()(baseHandlerFunc(difficultycontroller.Add, ""))).Methods("POST")
	router.Handle("/difficulty", middleware.EnsureValidToken()(baseHandlerFunc(difficultycontroller.GetAll, ""))).Methods("GET")
}

func baseHandlerFunc(next func(http.ResponseWriter, *http.Request), scope string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

		claims := token.CustomClaims.(*middleware.CustomClaims)

		if scope != "" && !claims.HasScope(scope) {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(`{"message":"Insufficient scope."}`))
			return
		}

		next(w, r)
	})
}

func setupHealthController(router *mux.Router) {
	router.Handle("/health", middleware.EnsureValidToken()(baseHandlerFunc(health, "")))
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("api is alive and well"))
}
