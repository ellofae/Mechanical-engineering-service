package auth

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	First_name string             `json:"first_name" bson:"first_name" validate:"required,lte=30"`
	Last_name  string             `json:"last_name" bson:"last_name" validate:"required,lte=30`
	Phone      string             `json:"phone" bson:"phone" validate:"required,lte=15`
	Password   string             `json:"password" bson:"password" validate:"required,lte=20`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
}

type SingInModel struct {
	Phone    string `json:"phone" bson:"phone" validate:"required,lte=15`
	Password string `json:"password" bson:"password" validate:"required,lte=20`
}
