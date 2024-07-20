package controller

import (
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"password-bank-golang/api/controller/model"
	"password-bank-golang/config"
	"password-bank-golang/service"
	"time"
)

type UserController struct {
	service service.IUserService
}

func (controller *UserController) Login(rw http.ResponseWriter, req *http.Request) {
	user, err := decodeUser(req)
	if err != nil {
		//I need to check all functions err handling mechanisms
		log.Printf("Error: %s", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = controller.service.LoginUser(user.Email, user.Password)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := createJWTToken(user.Email)
	if err != nil {
		log.Println("Error: Could not create token")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	writeTokenResponse(rw, token)
}

func (controller *UserController) Register(rw http.ResponseWriter, req *http.Request) {
	user, err := decodeUser(req)
	if err != nil {
		log.Printf("Error: %s", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = controller.service.CreateUser(user.Email, user.Password)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := createJWTToken(user.Email)
	if err != nil {
		log.Println("Error: Could not create token")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	writeTokenResponse(rw, token)
}

func writeTokenResponse(w http.ResponseWriter, token string) {
	err := json.NewEncoder(w).Encode(model.TokenResponse{
		Token: token,
	})
	w.WriteHeader(http.StatusOK)

	if err != nil {
		log.Println("Error: Could not parse token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func decodeUser(req *http.Request) (model.UserRequest, error) {
	var user model.UserRequest
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil || user.Email == "" || user.Password == "" {
		return user, errors.New("invalid credentials")
	}
	return user, nil
}

func createJWTToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Minute * 45).Unix(),
	})

	tokenStr, err := token.SignedString(config.SecretKey)

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func NewUserController(service service.IUserService) *UserController {
	return &UserController{service: service}
}
