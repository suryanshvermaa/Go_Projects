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
	dbHost := os.Getenv("DB_HOST");
	dbPort := os.Getenv("DB_PORT");
	dbUser := os.Getenv("DB_USER");
	dbPassword := os.Getenv("DB_PASSWORD");
	dbName := os.Getenv("DB_NAME");

	dsn := fmt.SPrintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort
	)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error in db connection:", err)
	}
	db = DB
}

func GetDB() *gorm.DB {
	return db
}