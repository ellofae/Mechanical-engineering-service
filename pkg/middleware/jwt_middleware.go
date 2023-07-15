package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtMiddleware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/joho/godotenv"
)

type Token struct {
	Expiry   int64
	IssuedAt int64
}

func GenerateToken() (string, error) {
	godotenv.Load(".env")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"expiry":   time.Now().Add(time.Minute * 30).Unix(),
		"issuedAt": time.Now().Unix(),
	})

	signedString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", fmt.Errorf("unable to create a signed JWT, error: %w", err)
	}

	return signedString, nil
}

func ParseToken(tokenString string) (*Token, error) {
	godotenv.Load(".env")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, fmt.Errorf("unable to parse token: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &Token{
			Expiry:   int64(claims["expiry"].(float64)),
			IssuedAt: int64(claims["issuedAt"].(float64)),
		}, nil
	} else {
		return nil, err
	}
}

func UserIdentity() func(ctx *fiber.Ctx) error {
	godotenv.Load(".env")

	config := jwtMiddleware.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET_KEY")),
		ErrorHandler: jwtError,
	}

	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": true,
		"msg":   err.Error(),
	})
}
