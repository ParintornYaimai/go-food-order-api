package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	Name       string             `bson:"name" json:"name" validate:"required"`
	Category   string             `bson:"category" json:"category" validate:"required"`
	Start_Date *time.Time         `bson:"start_date" json:"start_date,omitempty"`
	End_Date   *time.Time         `bson:"end_date" json:"end_date,omitempty"`
	Created_at *time.Time         `bson:"created_at" json:"created_at,omitempty"`
	Updated_at *time.Time         `bson:"updated_at" json:"updated_at,omitempty"`
	Menu_id    string             `bson:"menu_id" json:"menu_id" validate:"required"`
}
