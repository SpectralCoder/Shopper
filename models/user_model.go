package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is the model that governs all notes objects retrived or inserted into the DB
type User struct {
	ID              primitive.ObjectID `bson:"_id"`
	Name            *string            `json:"name" validate:"required,min=2,max=100"`
	Password        *string            `json:"Password" validate:"required,min=6"`
	Email           *string            `json:"email" validate:"email,required"`
	Phone           *int               `json:"phone" validate:"required"`
	Nid_no          *int               `json:"nid_no"`
	Token           *string            `json:"token"`
	Organization_id primitive.ObjectID `json:"org_id"`
	Refresh_token   *string            `json:"refresh_token"`
	Created_at      time.Time          `json:"created_at"`
	Updated_at      time.Time          `json:"updated_at"`
	User_id         string             `json:"user_id"`
}

type UserResponse struct {
	Name  *string `json:"name" validate:"required,min=2,max=100"`
	Email *string `json:"email" validate:"email,required"`
}
