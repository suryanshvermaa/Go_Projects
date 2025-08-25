package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID
	First_Name    *string   `json:"first_name" validate:"required, min=2, max=100"`
	Last_Name     *string   `json:"last_name" validate:"required, min=2, max=100"`
	Password      *string   `json:"-" validate:"required min=6"`
	Email         *string   `json:"email" validate:"required"`
	Phone         *string   `json:"phone" validate:"required"`
	Token         *string   `json:"token"`
	User_Type     *string   `json:"user_type" validate:"required, eq=ADMIN|eq=USER"`
	Refresh_Token *string   `json:"refresh_token"`
	Created_At    time.Time `json:"created_at"`
	Updated_At    time.Time `json:"updated_at"`
	User_Id       *string   `json:"user_id"`
}
