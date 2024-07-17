package controller

import (
	"fmt"
	"log"
	"net/http"
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

}

func (controller *PasswordController) GetAllPasswords(rw http.ResponseWriter, req *http.Request) {

}

func (controller *PasswordController) CreatePassword(rw http.ResponseWriter, req *http.Request) {

}

func (controller *PasswordController) UpdatePassword(rw http.ResponseWriter, req *http.Request) {

}

func (controller *PasswordController) DeletePassword(rw http.ResponseWriter, req *http.Request) {

}
