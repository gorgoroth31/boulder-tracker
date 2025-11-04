package middleware

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	Sub string `json:"sub"`
	jwt.StandardClaims
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		claims, err := verifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}

		r.Header.Set("principal", claims.Sub)
		next.ServeHTTP(w, r)
	})
}

func verifyToken(tokenString string) (*CustomClaims, error) {
	jwks, err := keyfunc.Get(os.Getenv("LOGTO_PUBLIC_KEY_LOCATION"), keyfunc.Options{})

	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, jwks.Keyfunc)

	if err != nil || !token.Valid {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims, nil
	}

	return nil, errors.New("something went wrong in token verification")
}
