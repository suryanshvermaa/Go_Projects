package models

import (
	"github.com/Go_Projects/golang-MySQL/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func (b *Book) GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func (b *Book) GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", Id).Find(&book)
	return &book, db
}

func (b *Book) DeleteBookById(Id int64) *Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return &book
}
