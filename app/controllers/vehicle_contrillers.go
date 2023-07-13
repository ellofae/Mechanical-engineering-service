package controllers

import (
	"log"
	"time"

	"github.com/ellofae/Mechanical-engineering-service/app/models"
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

/*
Create request test:
	{
    "vehicle_price": "NOT STATED",
    "category": "R35 GT-R",
    "title": "TS TSGold2022 Track Edition – Tokyo",
    "vehicle_status": "AVAILABLE",
    "model": "R35",
    "model_description": "Based on Nissan GT-R Track Edition engineered NISMO.TOP SECRET Full Bodykit with Special Paint “TS GOLD 2022”.TOKYO AUTOSALON 2022 Show Car.",
    "model_characteristics": [
        {
            "year": 1992,
            "mileage": 94.000,
            "engine": "VR38DETT",
            "engine_spec": "TOP SECRET Stage1",
            "suspensions": "TBA",
            "bodykit": "TopSecret MY17 Full Kit",
            "remarks": "Base: Track Edition engineered by NISMO / 2022 Tokyo AutoSalon Show Car"
        }
    ]
}
*/

func CreateVehicle(c *fiber.Ctx) error {
	vehicle := &models.Vehicle{}

	err := c.BodyParser(vehicle)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	vehicle.ID = uuid.New()
	vehicle.Created_at = time.Now()
	vehicle.Vehicle_status = "AVAILABE"

	log.Printf("test: ", vehicle)

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
	vehicle := &models.Vehicle{}

	err := c.BodyParser(vehicle)
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
	vehicle := &models.Service{}

	err := c.BodyParser(vehicle)
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
