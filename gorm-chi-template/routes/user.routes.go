package routes

import (
	"github.com/anshvermaa/go-chi-gorm/controllers"
	"github.com/go-chi/chi"
)

func UserRoutes() *chi.Mux {
	userRouter := chi.NewRouter()
	userRouter.Get("/getUser", controllers.GetUsers)
	return userRouter
}
