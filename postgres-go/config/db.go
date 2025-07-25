package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error in loading .env")
	}
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error to connect db")
	}
	DB = db
	fmt.Println("database connected!")
}

func GetDB() *gorm.DB {
	return DB
}
