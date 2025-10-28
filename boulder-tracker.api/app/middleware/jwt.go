package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v5"
)

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
		name := claims.(jwt.MapClaims)["name"].(string)
		role := claims.(jwt.MapClaims)["role"].(string)

		r.Header.Set("name", name)
		r.Header.Set("role", role)

		next.ServeHTTP(w, r)
	})
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	// hier muss ich über die API von logto den public key bekommen
	// https://bump.sh/logto/doc/logto-management-api/authentication
	// den Public key muss ich irgendwie benutzen um den Fehler rauszuschmeißen

	jwks, err := keyfunc.Get(os.Getenv("LOGTO_PUBLIC_KEY_LOCATION"), keyfunc.Options{})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	token, err := jwt.Parse(tokenString, jwks.Keyfunc)

	// claims := jwt.MapClaims{}

	/*token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return token, nil
	})*/

	return token.Claims, err
}
