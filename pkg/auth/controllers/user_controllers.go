package controllers

import (
	"context"
	"time"

	"github.com/ellofae/Mechanical-engineering-service/pkg/auth"
	"github.com/ellofae/Mechanical-engineering-service/pkg/middleware"
	"github.com/ellofae/Mechanical-engineering-service/platform/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SingUp(c *fiber.Ctx) error {
	user := &auth.User{}

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
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

	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()

	db.Collection.InsertOne(db.Ctx, user)

	return c.JSON(fiber.Map{
		"error": false,
		"user":  user,
	})
}

func SignIn(c *fiber.Ctx) error {
	loginData := &auth.SingInModel{}

	err := c.BodyParser(loginData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
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

	var result bson.M
	err = db.Collection.FindOne(context.Background(), bson.M{"phone": loginData.Phone, "password": loginData.Password}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	user := auth.User{}

	bsonBytes, _ := bson.Marshal(result)
	bson.Unmarshal(bsonBytes, &user)

	token, err := middleware.GenerateToken()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"user":  user,
		"token": token,
	})
}

func GetUser(c *fiber.Ctx) error {
	params := c.Queries()
	id := params["id"]

	db, err := database.OpenMongoDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	var result bson.M
	err = db.Collection.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	user := auth.User{}

	bsonBytes, _ := bson.Marshal(result)
	bson.Unmarshal(bsonBytes, &user)

	return c.JSON(fiber.Map{
		"error": false,
		"user":  user,
	})
}
