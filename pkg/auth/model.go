package auth

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	First_name string             `bson:"first_name"`
	Last_name  string             `bson:"last_name"`
	Email      string             `bson:"email"`
	Phone      string             `bson:"phone"`
	Password   string             `bson:"password"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
}

type SingInModel struct {
	Phone    string `bson:"phone"`
	Password string `bson:"password"`
}
