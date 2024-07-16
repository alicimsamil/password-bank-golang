package controller

import (
	"github.com/gorilla/mux"
	"net/http"
	"password-bank-golang/controller/middleware"
)

type PasswordBankController struct {
}

func (controller PasswordBankController) greetings(rw http.ResponseWriter, req *http.Request) {

}

func (controller PasswordBankController) getPasswordById(rw http.ResponseWriter, req *http.Request) {

}

func (controller PasswordBankController) getAllPasswords(rw http.ResponseWriter, req *http.Request) {

}

func (controller PasswordBankController) createPassword(rw http.ResponseWriter, req *http.Request) {

}

func (controller PasswordBankController) updatePassword(rw http.ResponseWriter, req *http.Request) {

}

func (controller PasswordBankController) deletePassword(rw http.ResponseWriter, req *http.Request) {

}

func (controller PasswordBankController) InitRouter() *mux.Router {

	middleWare := func(handler http.HandlerFunc) http.HandlerFunc {
		return middleware.LoggingMiddleWare(middleware.AuthMiddleWare(handler))
	}

	router := mux.NewRouter()
	router.Path("/").HandlerFunc(middleWare(controller.greetings)).Methods(http.MethodGet)
	router.Path("/password/all").HandlerFunc(middleWare(controller.getAllPasswords)).Methods(http.MethodGet)
	router.Path("/password/{id}").HandlerFunc(middleWare(controller.getPasswordById)).Methods(http.MethodGet)
	router.Path("/password").HandlerFunc(middleWare(controller.createPassword)).Methods(http.MethodPost)
	router.Path("/password/edit").HandlerFunc(middleWare(controller.updatePassword)).Methods(http.MethodPut)
	router.Path("/password/{id}").HandlerFunc(middleWare(controller.deletePassword)).Methods(http.MethodDelete)

	return router
}
