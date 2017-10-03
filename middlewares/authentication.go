package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

func Authentication(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		authHeader := request.Header.Get("Authorization")
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			/** @TODO When user control be implemented, get the token by user. */
			return []byte("secret"), nil
		})

		if err == nil && token.Valid {
			return
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			http.Error(writer, ve.Error(), http.StatusUnauthorized)
			return
		}

		http.Error(writer, "Unknown", http.StatusUnauthorized)
		return
	})
}
