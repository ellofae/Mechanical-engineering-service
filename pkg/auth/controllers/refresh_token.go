package controllers

import (
	"context"
	"errors"
	"time"

	"github.com/ellofae/Mechanical-engineering-service/pkg/auth"
	"github.com/ellofae/Mechanical-engineering-service/pkg/middleware"
	"github.com/ellofae/Mechanical-engineering-service/platform/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserID(c *fiber.Ctx) (*primitive.ObjectID, error) {
	var user_id string
	if c.Locals("user_id") == nil {
		return nil, errors.New("unable to get user's uuid.UUID from the local's state")
	}

	user_id = c.Locals("user_id").(string)
	objectId, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return nil, err
	}

	return &objectId, nil
}

func ParseMongoObject(ctx context.Context, objectId primitive.ObjectID, db *database.MongoDB) (*auth.User, error) {
	var result bson.M
	err := db.Collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}

	user := auth.User{}
	bsonBytes, _ := bson.Marshal(result)
	bson.Unmarshal(bsonBytes, &user)

	return &user, nil
}

func CheckToken(user *auth.User) error {
	if user.Refresh_token == "" {
		return errors.New("no refresh token available in database")
	}

	token_claims, err := middleware.ParseToken(user.Refresh_token)
	if err != nil {
		return err
	}

	if token_claims.Expiry < time.Now().Unix() {
		return errors.New("token expired")
	}

	return nil
}

func RefreshToken(c *fiber.Ctx) error {
	objectId, err := GetUserID(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := database.OpenMongoDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user, err := ParseMongoObject(ctx, *objectId, db)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	err = CheckToken(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	tokens, err := middleware.GenerateTokens(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	cookie := fiber.Cookie{
		Name:     "access_token",
		Value:    tokens.AccessToken,
		Expires:  time.Now().Add(time.Hour * 3),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	err = AddRefreshToken(ctx, db, tokens.RefreshToken, user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":         false,
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
	})
}
