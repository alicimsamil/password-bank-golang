package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"password-bank-golang/api/controller/model"
	"password-bank-golang/config"
	"password-bank-golang/service"
	"strings"
)

type PasswordController struct {
	service service.IPasswordService
}

func (controller *PasswordController) Greetings(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "Hi! Welcome to my password API. I hope you enjoy it!")
}

func (controller *PasswordController) GetPasswordById(rw http.ResponseWriter, req *http.Request) {
	email, err := getUserEmail(req)
	if err != nil {
		http.Error(rw, "Invalid user email", http.StatusBadRequest)
		return
	}

	var args = mux.Vars(req)
	password, err := controller.service.GetPasswordById(args["id"], email)
	if err != nil {
		http.Error(rw, "Password not found", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(rw).Encode(
		model.GetPasswordResponse{
			Id:          password.Id,
			Password:    password.Password,
			Type:        password.Type,
			Account:     password.Account,
			ServiceName: password.ServiceName,
			Notes:       password.Notes,
			Date:        password.Date})

	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (controller *PasswordController) GetAllPasswords(rw http.ResponseWriter, req *http.Request) {
	email, err := getUserEmail(req)
	if err != nil {
		http.Error(rw, "Invalid user email", http.StatusBadRequest)
		return
	}

	passwords, err := controller.service.GetAllPasswords(email)
	if err != nil {
		http.Error(rw, "Passwords not found", http.StatusNotFound)
		return
	}
	var allPasswords []model.GetPasswordResponse
	for _, element := range passwords {
		allPasswords = append(allPasswords, model.GetPasswordResponse{
			Id:          element.Id,
			Password:    element.Password,
			Type:        element.Type,
			Account:     element.Account,
			ServiceName: element.ServiceName,
			Notes:       element.Notes,
			Date:        element.Date})
	}

	if err := json.NewEncoder(rw).Encode(allPasswords); err != nil {
		http.Error(rw, "Failed to encode passwords", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (controller *PasswordController) SavePassword(rw http.ResponseWriter, req *http.Request) {
	email, err := getUserEmail(req)
	if err != nil {
		http.Error(rw, "Invalid user email", http.StatusBadRequest)
		return
	}

	var passReq model.PasswordRequest
	if err := json.NewDecoder(req.Body).Decode(&passReq); err != nil {
		http.Error(rw, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := controller.service.InsertPassword(passReq, email); err != nil {
		http.Error(rw, "Failed to save password", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	fmt.Fprintln(rw, "Your password saved.")
}

func (controller *PasswordController) UpdatePassword(rw http.ResponseWriter, req *http.Request) {
	email, err := getUserEmail(req)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var passReq model.PasswordRequest
	if err := json.NewDecoder(req.Body).Decode(&passReq); err != nil {
		http.Error(rw, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := controller.service.UpdatePassword(passReq, email); err != nil {
		http.Error(rw, "Failed to update password", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	fmt.Fprintln(rw, "Your password updated.")
}

func (controller *PasswordController) DeletePassword(rw http.ResponseWriter, req *http.Request) {
	email, err := getUserEmail(req)
	if err != nil {
		http.Error(rw, "Invalid user email", http.StatusBadRequest)
		return
	}

	args := mux.Vars(req)
	if err := controller.service.DeletePassword(args["id"], email); err != nil {
		http.Error(rw, "Password not found", http.StatusNotFound)
		return
	}

	rw.WriteHeader(http.StatusOK)
	fmt.Fprintln(rw, "Password deleted.")
}

func getUserEmail(r *http.Request) (string, error) {
	auth := r.Header.Get("Authorization")

	parsed := strings.Split(auth, " ")[1]

	token, err := jwt.Parse(parsed, func(token *jwt.Token) (interface{}, error) {
		return config.SecretKey, nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("token didn't parsed")
	}

	email, ok := claims["email"].(string)
	if !ok || email == "" {
		return "", errors.New("email not found in token")
	}

	return email, nil
}

func NewPasswordController(service service.IPasswordService) *PasswordController {
	return &PasswordController{service: service}
}
