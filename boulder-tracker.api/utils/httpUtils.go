package utils

import (
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

func GetPrincipalId(r *http.Request) string {
	token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

	return token.RegisteredClaims.Subject
}
