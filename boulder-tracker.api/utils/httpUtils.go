package utils

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func GetPrincipalId(r *http.Request) string {
	//token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

	return ""
}

func GetToken(name string) (string, error) {
	signingKey := []byte("keymaker")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
		"role": "redpill",
	})
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

func VerifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte("keymaker")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}
