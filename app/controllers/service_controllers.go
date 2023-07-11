package controllers

import (
	"github.com/ellofae/Mechanical-engineering-service/app/models"
	"github.com/ellofae/Mechanical-engineering-service/platform/database"
	"github.com/ellofae/Mechanical-engineering-service/pkg/utils"

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

func UpdateBooks(c *fiber.Ctx) error {
	service := &models.Service{}

	if err := c.BodyParser(service); err != nil {
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

	foundService, err := db.GetService(service.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg": err.Error()
		})
	}

	serivce.Updated_at = time.Now()

	validate := utils.NewValidator()
	
	if err := validate.Struct(service); err != nil {
		return c.Status(fiber.StatusBadRequst).JSON(fiber.Map{
			"error": true,
			"msg": utils.ValidatorErrors(err)
		})
	}

	if err := db.UpdateService(foundService.ID, service); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}

func DeleteService(c *fiber.Ctx) error {
	service := &models.Service{}

	if err := c.BodyParser(service); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg": err.Error(),
		})
	}

	validate := utils.NewValidator()
	
	if err := validate.StructPartial(service, "id"); err != nil {
		return c.Status(fiber.StatusBadRequst).JSON(fiber.Map{
			"error": true,
			"msg": utils.ValidatorErrors(err)
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg": err.Error(),
		})
	}

	foundService, err := db.GetService(service.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg": err.Error()
		})
	}

	if err := db.DeleteService(foundService.ID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}