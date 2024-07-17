package controller

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
)

type PasswordController struct {
}

func (controller *PasswordController) Greetings(rw http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintln(rw, "Hi! Welcome to my password API. I hope you enjoy it!")
	if err != nil {
		log.Println("Error writing response: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

func (controller *PasswordController) GetPasswordById(rw http.ResponseWriter, req *http.Request) {

	email, err := getUserEmail(req)
	//I need to check all functions err handling mechanisms
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (controller *PasswordController) GetAllPasswords(rw http.ResponseWriter, req *http.Request) {
	email, err := getUserEmail(req)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (controller *PasswordController) CreatePassword(rw http.ResponseWriter, req *http.Request) {
	email, err := getUserEmail(req)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (controller *PasswordController) UpdatePassword(rw http.ResponseWriter, req *http.Request) {
	email, err := getUserEmail(req)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (controller *PasswordController) DeletePassword(rw http.ResponseWriter, req *http.Request) {
	email, err := getUserEmail(req)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
}

// I am not sure about this function
func getUserEmail(r *http.Request) (string, error) {
	auth := r.Header.Get("Authorization")

	parsed := strings.Split(auth, " ")[1]

	token, err := jwt.Parse(parsed, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return "", errors.New("token didn't parsed")
	}

	email := claims["email"].(string)

	if email == "" {
		return "", errors.New("email not found in token")
	}

	return email, nil
}
