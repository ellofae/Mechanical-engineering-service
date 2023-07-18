package configs

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func NewConfig() fiber.Config {
	godotenv.Load(".env")

	engine := html.New("./views", ".html")

	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	writeTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_WRITE_TIMEOUT"))
	idleTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_IDLE_TIMEOUT"))

	return fiber.Config{
		AppName:      "Mechanical engineering service",
		Views:        engine,
		ReadTimeout:  time.Second * time.Duration(readTimeoutSecondsCount),
		WriteTimeout: time.Second * time.Duration(writeTimeoutSecondsCount),
		IdleTimeout:  time.Second * time.Duration(idleTimeoutSecondsCount),
	}
}
