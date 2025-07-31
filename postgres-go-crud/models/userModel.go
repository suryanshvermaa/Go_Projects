package models

import (
	"time"

	"github.com/Go_Projects/postgres-go-crud/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Contact struct {
	Email   string `json:"email" gorm:"unique;not null"` // allowing for null values also so this is pointer
	PhoneNo string `json:"phone_no" gorm:"unique;not null"`
}

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey; autoIncrement"`
	Name      string `json:"name"`
	Contact   `gorm:"embedded;embeddedPrefix:contact_"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"` // it is the specific field which should be as it is for GORM
	UpdatedAt time.Time `json:"updated_at"` // it is the specific field which should be as it is for GORM
}

func UserModelInit() {
	DB := config.GetDB()
	db = DB
	db.AutoMigrate(&User{})
}
func (u *User) CreateUser() (*User, error) {
	db := config.GetDB()
	result := db.Create(&u)
	if u.ID == 0 {
		return nil, result.Error
	}
	return u, nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	result := db.First(&user).Where("Email=?", email)
	if user.ID == 0 {
		return nil, result.Error
	}
	return &user, nil
}
