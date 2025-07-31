package models

import (
	"time"

	"github.com/Go_Projects/postgres-go-crud/config"
)

type Contact struct {
	Email   *string `json:"email" gorm:"unique; not null"` // allowing for null values also so this is pointer
	PhoneNo string  `json:"phone_no" gorm:"unique; not null"`
}

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey; autoIncreament"`
	Name      string    `json:"name"`
	Contact   Contact   `gorm:"embedded"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"` // it is the specific field which should be as it is for GORM
	UpdatedAt time.Time `json:"updated-at"` // it is the specific field which should be as it is for GORM
}

func init() {
	db := config.GetDB()
	db.AutoMigrate(&User{})
}
func (u *User) CreateUser(name string, email string, password string, phone string) *User {
	db := config.GetDB()
	db.Create(&u)
	return u
}
