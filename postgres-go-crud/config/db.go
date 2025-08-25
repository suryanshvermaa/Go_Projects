package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDb() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("err in loading .env")
	}
	dsn := os.Getenv("DSN")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error in db connection:", err)
	}
	db = DB
}

func GetDB() *gorm.DB {
	return db
}
