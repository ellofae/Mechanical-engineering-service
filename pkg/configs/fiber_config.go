package configs

import (
	"os"
	"strconv"
	"time"
	
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func NewConfig() fiber.Config {
	godotenv.Load(".env")

	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	writeTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_WRITE_TIMEOUT"))
	idleTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_IDLE_TIMEOUT"))

	return fiber.Config{
		AppName: "Mechanical engineering service",
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
		WriteTimeout: time.Second * time.Duration(writeTimeoutSecondsCount),
		IdleTimeout: time.Second * time.Duration(idleTimeoutSecondsCount),
	}
}