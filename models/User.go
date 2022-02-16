package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID      `json:"_id" bson:"_id,omitempty"`
	First_name     *string                 `json:"first_name" validate:"required,min=2,max=100"`
	Last_name      *string                 `json:"last_name" validate:"required, min=2, max=100"`
	Email          *string                 `json:"email" validate:"required"`
	Phone          *string                 `json:"phone" validate:"required"`
	User_type      *string                 `json:"user_type" validate:"required, eq=ADMIN|eq=USER"`
	Created_at     time.Time			   `json:"created_at"`
	Updated_at	   time.Time               `json:"updated_at"`
	User_id		   *string 				   `json:"user_id"`
}