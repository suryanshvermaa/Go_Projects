package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Table struct {
	ID               primitive.ObjectID `bson:"_id"`
	Number_Of_Guests *int               `json:"number_of_guests" validate:"required"`
	Table_Number     *int               `json:"table_number" validate:"required"`
	Created_At       time.Time          `json:"created_at"`
	Updated_At       time.Time          `json:"updated_at"`
	Table_Id         string             `json:"table_id"`
}
