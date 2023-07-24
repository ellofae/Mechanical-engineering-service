package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/ellofae/Mechanical-engineering-service/pkg/auth"
	"github.com/ellofae/Mechanical-engineering-service/pkg/auth/hashing"
	"github.com/ellofae/Mechanical-engineering-service/pkg/middleware"
	"github.com/ellofae/Mechanical-engineering-service/pkg/utils"
	"github.com/ellofae/Mechanical-engineering-service/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func checkUserExists(db *database.MongoDB, phone string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var result bson.M
	err := db.Collection.FindOne(ctx, bson.M{"phone": phone}).Decode(&result)
	if err != nil {
		return err
	}

	return nil
}

func SingUp(c *fiber.Ctx) error {
	user := &auth.User{}
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if user.Phone == "" || user.Password == "" || user.First_name == "" || user.Last_name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "All fields must contain data",
			"req":   user,
		})
	}

	validate := utils.NewValidator()

	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()

	err = validate.Struct(user)
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

	err = checkUserExists(db, user.Phone)
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "User with such phone number already exists",
		})
	}

	passwordHashed, err := hashing.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	user.Password = passwordHashed

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = db.Collection.InsertOne(ctx, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"id":         user.ID,
		"first_name": user.First_name,
		"last_name":  user.Last_name,
		"phone":      user.Phone,
	})
}

func AddRefreshToken(ctx context.Context, db *database.MongoDB, refresh_token string, ID primitive.ObjectID) error {
	_, err := middleware.ParseToken(refresh_token)
	if err != nil {
		return err
	}

	filter := bson.D{{"_id", ID}}
	update := bson.D{{"$set", bson.D{{"refresh_token", refresh_token}}}}

	result, err := db.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	fmt.Printf("Documents updated: %v\nNew document: %v\n", result.ModifiedCount, result)

	return nil
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

	validate := utils.NewValidator()

	err = validate.Struct(loginData)
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

	var result bson.M
	err = db.Collection.FindOne(ctx, bson.M{"phone": loginData.Phone}).Decode(&result)
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

	checkPassword := hashing.CheckPasswordHash(loginData.Password, user.Password)
	if !checkPassword {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Incorrect password passed",
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
	c.Locals("user_id", user.ID)

	err = AddRefreshToken(ctx, db, tokens.RefreshToken, user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"id":           user.ID,
		"first_name":   user.First_name,
		"last_name":    user.Last_name,
		"access_token": tokens.AccessToken,
	})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var result bson.M
	err = db.Collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&result)
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

func GetUsers(c *fiber.Ctx) error {
	godotenv.Load(".env")
	db, err := database.OpenMongoDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := db.Collection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	defer cursor.Close(ctx)

	users := []auth.User{}
	for cursor.Next(ctx) {
		var user auth.User

		cursor.Decode(&user)
		users = append(users, user)
	}

	if err = cursor.Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"users": users,
	})
}
