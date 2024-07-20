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

	var args = mux.Vars(req)
	password, err := controller.service.GetPasswordById(args["id"], email)
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
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
		//I need to check all functions err handling mechanisms
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var allPasswords []model.GetPasswordResponse
	passwords, err := controller.service.GetAllPasswords(email)

	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

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

	err = json.NewEncoder(rw).Encode(allPasswords)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (controller *PasswordController) SavePassword(rw http.ResponseWriter, req *http.Request) {
	email, err := getUserEmail(req)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var passReq model.PasswordRequest

	err = json.NewDecoder(req.Body).Decode(&passReq)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = controller.service.InsertPassword(passReq, email)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(rw, "Your password saved.")
	rw.WriteHeader(http.StatusOK)
}

func (controller *PasswordController) UpdatePassword(rw http.ResponseWriter, req *http.Request) {
	email, err := getUserEmail(req)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var passReq model.PasswordRequest

	err = json.NewDecoder(req.Body).Decode(&passReq)

	err = controller.service.UpdatePassword(passReq, email)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(rw, "Your password updated.")
	rw.WriteHeader(http.StatusOK)
}

func (controller *PasswordController) DeletePassword(rw http.ResponseWriter, req *http.Request) {
	email, err := getUserEmail(req)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var args = mux.Vars(req)
	err = controller.service.DeletePassword(args["id"], email)
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
}

// I am not sure about this function
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

	email := claims["email"].(string)

	if email == "" {
		return "", errors.New("email not found in token")
	}

	return email, nil
}

func NewPasswordController(service service.IPasswordService) *PasswordController {
	return &PasswordController{service: service}
}
