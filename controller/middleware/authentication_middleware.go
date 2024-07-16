package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
)

func AuthMiddleWare(handler http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")
		if authHeader == "" {
			rw.WriteHeader(http.StatusUnauthorized)
			log.Printf("Unauthorized action! Uri: %s", req.RequestURI)
			return
		}

		token := strings.Split(authHeader, " ")

		if len(token) != 2 {
			rw.WriteHeader(http.StatusUnauthorized)
			log.Printf("Unauthorized action! Uri: %s", req.RequestURI)
			return
		}

		err := verifyToken(token[1])
		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			log.Println(err.Error())
			return
		}

		handler.ServeHTTP(rw, req)
	}
}

func verifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return token, nil
	})

	if err != nil {
		return err
	}

	if !parsedToken.Valid {
		return errors.New("invalid token")
	}

	return nil
}
