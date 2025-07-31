package main

import (
	"log"
	"net/http"
	"os"

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
	err = http.ListenAndServe(":"+port, router)
	log.Printf("Server running on port %s", port)
	if err != nil {
		log.Fatal("Error in running server", err)
	}
}
