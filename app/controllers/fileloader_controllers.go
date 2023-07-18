package controllers

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

const UPLOAD_FOLDER string = "public/uploads/"

func LoadingFileFields(c *fiber.Ctx) error {
	return c.Render("fileload", fiber.Map{})
}

func LoadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("upload")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	c.SaveFile(file, UPLOAD_FOLDER+file.Filename)
	return c.Render("fileload", fiber.Map{
		"Name": file.Filename,
	})
}

type FileForTemplate struct {
	Filename string
}

func AccessingFile(c *fiber.Ctx) error {
	entries, err := os.ReadDir(UPLOAD_FOLDER)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	fileSlice := []FileForTemplate{}
	for _, e := range entries {
		fileSlice = append(fileSlice, FileForTemplate{Filename: e.Name()})
	}

	return c.Render("filesave", fiber.Map{
		"Files": fileSlice,
	})
}

func AccessFile(c *fiber.Ctx) error {
	document := c.FormValue("document")

	entries, err := os.ReadDir(UPLOAD_FOLDER)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	flag := false
	for _, e := range entries {
		if e.Name() == document {
			flag = true
		}
	}

	if !flag {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "No file with such name exists in storage",
		})
	}

	return c.SendFile(UPLOAD_FOLDER + document)
}
