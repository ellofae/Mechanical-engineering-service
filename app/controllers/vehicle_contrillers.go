package controllers

import (
	"time"

	"github.com/ellofae/Mechanical-engineering-service/app/models"
	"github.com/ellofae/Mechanical-engineering-service/pkg/middleware"
	"github.com/ellofae/Mechanical-engineering-service/pkg/utils"
	"github.com/ellofae/Mechanical-engineering-service/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetVehicles(c *fiber.Ctx) error {
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	vehicles, err := db.GetVehicles()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":    true,
			"count":    0,
			"vehicles": nil,
			"msg":      err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":    false,
		"count":    len(vehicles),
		"vehicles": vehicles,
	})
}

func GetVehicle(c *fiber.Ctx) error {
	params := c.Queries()

	id, err := uuid.Parse(params["id"])
	if err != nil {
		return err
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	vehicle, err := db.GetVehicle(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"vehicle": nil,
			"msg":     err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"vehicle": vehicle,
	})
}

func CreateVehicle(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := middleware.ParseToken(cookie)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	expiry := token.Expiry

	if expiry < time.Now().Unix() {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"message": "token expired",
		})
	}

	vehicle := &models.Vehicle{}

	err = c.BodyParser(vehicle)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	vehicle.ID = uuid.New()
	vehicle.Created_at = time.Now()
	vehicle.Vehicle_status = "AVAILABE"

	validate := utils.NewValidator()

	err = validate.Struct(vehicle)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	err = db.CreateVehicle(vehicle)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}

func UpdateVehicle(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := middleware.ParseToken(cookie)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	expiry := token.Expiry

	if expiry < time.Now().Unix() {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"message": "token expired",
		})
	}

	vehicle := &models.Vehicle{}

	err = c.BodyParser(vehicle)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	vehicle.Updated_at = time.Now()

	validate := utils.NewValidator()
	err = validate.Struct(vehicle)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	foundVehicle, err := db.GetVehicle(vehicle.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	err = db.UpdateVehicle(foundVehicle.ID, vehicle)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}

func DeleteVehicle(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := middleware.ParseToken(cookie)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	expiry := token.Expiry

	if expiry < time.Now().Unix() {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"message": "token expired",
		})
	}

	vehicle := &models.Service{}

	err = c.BodyParser(vehicle)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	validate := utils.NewValidator()
	err = validate.StructPartial(vehicle, "id")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	foundVehicle, err := db.GetVehicle(vehicle.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	err = db.DeleteVehicle(foundVehicle.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
