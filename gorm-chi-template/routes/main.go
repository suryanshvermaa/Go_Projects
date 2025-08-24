package routes

import "github.com/go-chi/chi"

func SetupRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Mount("/user", UserRoutes())
	return router
}
