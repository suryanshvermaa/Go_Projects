package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Go_Projects/postgres-go-crud/config"
	"github.com/Go_Projects/postgres-go-crud/models"
	"github.com/Go_Projects/postgres-go-crud/routes"
	"github.com/Go_Projects/postgres-go-crud/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error in loading dotenv file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	router := mux.NewRouter()
	// db connection
	config.InitDb()
	models.UserModelInit()

	routes.InitUserRoutes(router)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.JsonResponse(w, 200, "server is running healty", nil)
	}).Methods("GET")
	log.Printf("Server running on port %s", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("Error in running server", err)
	}
}
