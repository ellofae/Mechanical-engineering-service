package utils

import (
	"log"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2"
)

func StartServerWithGracefulShutdown(a *fiber.App) {
	godotenv.Load(".env")

	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)

		<-sigint
		
		if err := a.Shutdown(); err != nil {
			log.Printf("Server is not shutting down! Reason: %w", err)
		}
		close(idleConnsClosed)
	}

	if err := a.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Printf("Server is not running! Reason: %w", err)
	}
	<-idleConnsClosed
}