package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name"`
	Category   string             `json:"category"`
	Start_Date *time.Time         `json:"start_date"`
	End_Date   *time.Time         `json:"end_time"`
	Created_At time.Time          `json:"created_at"`
	Updated_At time.Time          `json:"updated_at"`
	Menu_Id    *string            `json:"menu_id"`
}
