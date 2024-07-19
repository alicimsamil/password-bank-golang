package main

import (
	"log"
	"net/http"
	"password-bank-golang/api/controller"
	"password-bank-golang/api/router"
	"password-bank-golang/data/db"
	"password-bank-golang/data/repository"
	"password-bank-golang/service"
)

func main() {
	dbConn, err := db.GetDBConn()
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(dbConn)
	passRepo := repository.NewPasswordRepository(dbConn)

	userService := service.NewUserService(userRepo)
	passService := service.NewPasswordService(passRepo)

	userController := controller.NewUserController(userService)
	passController := controller.NewPasswordController(passService)

	mRouter := router.NewPasswordBankRouter(userController, passController)

	log.Fatal(http.ListenAndServe(":80", mRouter.InitRouter()))
}
