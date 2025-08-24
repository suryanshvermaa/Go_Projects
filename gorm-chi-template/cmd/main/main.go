package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/anshvermaa/go-chi-gorm/routes"
	"github.com/anshvermaa/go-chi-gorm/utils"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error to load .env file")
	}
	port := os.Getenv("PORT")
	router := chi.NewRouter()
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		utils.JsonResponse(w, 200, "server is healty", nil)
	})
	router.Mount("/", routes.SetupRoutes())

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	fmt.Printf("server is running on port: %s\n", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
