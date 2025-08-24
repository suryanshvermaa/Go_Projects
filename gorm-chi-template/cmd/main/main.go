package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	})
	fmt.Println("Hello")
}
