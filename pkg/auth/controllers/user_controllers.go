package controllers

import (
	"context"
	"time"

	"github.com/ellofae/Mechanical-engineering-service/pkg/auth"
	"github.com/ellofae/Mechanical-engineering-service/pkg/auth/hashing"
	"github.com/ellofae/Mechanical-engineering-service/pkg/middleware"
	"github.com/ellofae/Mechanical-engineering-service/pkg/utils"
	"github.com/ellofae/Mechanical-engineering-service/platform/database"
	"github.com/gofiber/fiber/v2"
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
	// Form data
	first_name := c.FormValue("firstName")
	last_name := c.FormValue("lastName")
	phone := c.FormValue("phone")
	password := c.FormValue("password")

	user := &auth.User{
		First_name: first_name,
		Last_name:  last_name,
		Phone:      phone,
		Password:   password,
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

	err := validate.Struct(user)
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

	return c.Render("greeting", fiber.Map{
		"First_name": user.First_name,
		"Last_name":  user.Last_name,
		"Phone":      user.Phone,
	})
}

func SignIn(c *fiber.Ctx) error {
	phone := c.FormValue("phone")
	password := c.FormValue("password")

	loginData := &auth.SingInModel{
		Phone:    phone,
		Password: password,
	}

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

	token, err := middleware.GenerateToken()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Render("login", fiber.Map{
		"First_name": user.First_name,
		"Last_name":  user.Last_name,
		"Token":      token,
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

func RegisterUser(c *fiber.Ctx) error {
	return c.Render("signup", fiber.Map{})
}

func LoginrUser(c *fiber.Ctx) error {
	return c.Render("signin", fiber.Map{})
}
