package routes

import (
	"net/http"

	"github.com/Go_Projects/postgres-go-crud/controllers"
	"github.com/Go_Projects/postgres-go-crud/middlewares"
	"github.com/gorilla/mux"
)

func InitUserRoutes(router *mux.Router) {
	router.HandleFunc("/api/user/signup", controllers.SignUp).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Login).Methods("POST")
	router.Handle("/api/user/profile", middlewares.AuthMiddleware(http.HandlerFunc(controllers.Profile))).Methods("GET")
}
