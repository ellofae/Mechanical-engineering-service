package controllers

import (
	"github.com/ellofae/Mechanical-engineering-service/app/models"
	"github.com/ellofae/Mechanical-engineering-service/platform/database"

	"time"
	"github.com/google/uuid"
	"github.com/gofiber/fiber/v2"
)

func GetServices(c *fiber.Ctx) error {
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg": err.Error(),
		})
	}

	services, err := db.GetServices()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg": err.Error(),
			"count": 0,
			"services": nil,
		})
	}

	return c.JSON(fiber.Map{
		"error": "false",
		"msg": nil,
		"count": len(services),
		"services": services,
	})
}

func CreateService(c *fiber.Ctx) error {
	service := models.Service{}

	if err := c.BodyParser(&service); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg": err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg": err.Error(),
		})
	}

	service.ID = uuid.New()
	service.Created_at = time.Now()
	service.Service_status = "AVAILABLE"

	if err := db.CreateService(&service); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg": nil,
		"service": service,
	})
}