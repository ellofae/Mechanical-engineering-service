package middleware

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/joho/godotenv"
)

type TokensResponse struct {
	AccessToken  string
	RefreshToken string
}

type Token struct {
	Expiry   int64
	IssuedAt int64
	UserID   string
	State    string
}

func GenerateTokens(user_id primitive.ObjectID) (*TokensResponse, error) {
	godotenv.Load(".env")

	access_token_payload := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"expiry":   time.Now().Add(time.Minute * 30).Unix(),
		"issuedAt": time.Now().Unix(),
		"user_id":  user_id,
		"state":    "access_token",
	})

	refresh_token_payload := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"expiry":   time.Now().Add(time.Hour * 24).Unix(),
		"issuedAt": time.Now().Unix(),
		"user_id":  user_id,
		"state":    "refresh_token",
	})

	access_signedString, err := access_token_payload.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return nil, fmt.Errorf("unable to create an access signed JWT, error: %w", err)
	}

	refresh_signedString, err := refresh_token_payload.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return nil, fmt.Errorf("unable to create a refresh signed JWT, error: %w", err)
	}

	return &TokensResponse{AccessToken: access_signedString, RefreshToken: refresh_signedString}, nil
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
			UserID:   claims["user_id"].(string),
			State:    claims["state"].(string),
		}, nil
	} else {
		return nil, err
	}
}

func AuthenticateUser(c *fiber.Ctx) error {
	var tokenString string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("access_token") != "" {
		tokenString = c.Cookies("access_token")
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
	}

	token_claims, err := ParseToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Incorrect access token"})
	}

	c.Locals("user_id", token_claims.UserID)

	expiry := token_claims.Expiry
	if expiry < time.Now().Unix() {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"message": "token expired",
		})
	}

	c.Locals("issuedAt", token_claims.IssuedAt)
	return c.Next()
}
