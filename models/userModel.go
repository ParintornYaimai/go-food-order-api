package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id"`
	First_name    *string            `bson:"first_name" json:"first_name" validate:"required,min=2,max=100"`
	Last_name     *string            `bson:"last_name" json:"last_name" validate:"required,min=2,max=100"`
	Password      *string            `bson:"password" json:"password" validate:"required,min=6"`
	Email         *string            `bson:"email" json:"email" validate:"email, required"`
	Avatar        *string            `bson:"avatar" json:"avatar"`
	Phone         *string            `bson:"phone" json:"phone" validate:"required"`
	Token         *string            `bson:"token" json:"token"`
	Refresh_Token *string            `bson:"refresh_Token" json:"refresh_Token"`
	Created_at    time.Time          `bson:"created_at" json:"created_at"`
	Updated_at    time.Time          `bson:"updated_at" json:"updated_at"`
	User_id       string             `bson:"user_id" json:"user_id"`
}
