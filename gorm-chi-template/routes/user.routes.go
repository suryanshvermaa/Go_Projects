package routes

import (
	"github.com/anshvermaa/go-chi-gorm/controllers"
	"github.com/anshvermaa/go-chi-gorm/middlewares"
	"github.com/go-chi/chi"
)

func UserRoutes() *chi.Mux {
	userRouter := chi.NewRouter()
	userRouter.Use(middlewares.AuthMiddleware)
	userRouter.Get("/getUser", controllers.GetUsers)
	return userRouter
}
