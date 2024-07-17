package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"password-bank-golang/api/controller"
	"password-bank-golang/api/middleware"
)

type PasswordBankRouter struct {
	passController *controller.PasswordController
	userController *controller.UserController
}

func (c *PasswordBankRouter) InitRouter() *mux.Router {

	middleWare := func(handler http.HandlerFunc) http.HandlerFunc {
		return middleware.LoggingMiddleWare(middleware.AuthMiddleWare(handler))
	}

	router := mux.NewRouter()

	//Password operations
	router.Path("/").HandlerFunc(middleWare(c.passController.greetings)).Methods(http.MethodGet)
	router.Path("/password/all").HandlerFunc(middleWare(c.passController.getAllPasswords)).Methods(http.MethodGet)
	router.Path("/password/{id}").HandlerFunc(middleWare(c.passController.getPasswordById)).Methods(http.MethodGet)
	router.Path("/password").HandlerFunc(middleWare(c.passController.createPassword)).Methods(http.MethodPost)
	router.Path("/password/edit").HandlerFunc(middleWare(c.passController.updatePassword)).Methods(http.MethodPut)
	router.Path("/password/{id}").HandlerFunc(middleWare(c.passController.deletePassword)).Methods(http.MethodDelete)

	//User operations
	router.Path("/login").HandlerFunc(middleWare(c.userController.login)).Methods(http.MethodPost)
	router.Path("/login").HandlerFunc(middleWare(c.userController.register)).Methods(http.MethodPost)

	return router
}
