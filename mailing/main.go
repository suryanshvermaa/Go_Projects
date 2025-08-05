package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error in loading .env")
	}
	mail := os.Getenv("MY_EMAIL")
	password := os.Getenv("MY_PASSWORD")
	SendMail(mail, password, Mail)
	fmt.Println("Hello")
}
