package controller

import (
	"fmt"
	"log"
	"net/http"
)

type PasswordController struct {
}

func (controller *PasswordController) greetings(rw http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintln(rw, "Hi! Welcome to my password API. I hope you enjoy it!")
	if err != nil {
		log.Println("Error writing response: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

func (controller *PasswordController) getPasswordById(rw http.ResponseWriter, req *http.Request) {

}

func (controller *PasswordController) getAllPasswords(rw http.ResponseWriter, req *http.Request) {

}

func (controller *PasswordController) createPassword(rw http.ResponseWriter, req *http.Request) {

}

func (controller *PasswordController) updatePassword(rw http.ResponseWriter, req *http.Request) {

}

func (controller *PasswordController) deletePassword(rw http.ResponseWriter, req *http.Request) {

}
