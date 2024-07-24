package routes

import (
	"github.com/gorilla/mux"

	"github.com/uday778/LibraryManagementSystem-with-gorillaMux/controllers"

)

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/login",controllers.Login)
	router.HandleFunc("/signup",controllers.SignUp)
}