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
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not defined")
	}
	router := chi.NewRouter()

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	fmt.Printf("server is running on port: %s\n", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
