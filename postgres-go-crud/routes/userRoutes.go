package routes

import (
	"github.com/Go_Projects/postgres-go-crud/controllers"
	"github.com/gorilla/mux"
)

func InitUserRoutes(router *mux.Router) {
	router.HandleFunc("/api/user/signup", controllers.SignUp).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Login).Methods("POST")
}
