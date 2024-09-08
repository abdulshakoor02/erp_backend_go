package auth

import (
	// "encoding/json"
	// "fmt"
	//
	// "github.com/abdul/erp_backend/database/dbAdapter"

	"fmt"

	"github.com/abdul/erp_backend/logger"

	"github.com/abdul/erp_backend/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var log = logger.Logger

type AuthParam struct {
	Token string `json:"token"`
}

func AuthHandler(c *fiber.Ctx) error {
	t := new(AuthParam)
	if err := c.BodyParser(t); err != nil {
		return err
	}
	// to do handle token verification here
	secretKey := []byte(config.SECRET_KEY)
	token, err := jwt.Parse(t.Token, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		log.Err(err).Msgf("failed to sign token : %v", err)
		return c.Status(fiber.StatusUnauthorized).SendString("invalid Request")
	}

	parsedClaim := token.Claims.(jwt.MapClaims)
	if parsedClaim["email"] == "test@test.com" && parsedClaim["password"] == "test@123" {
		fmt.Println("user authorised")
		return c.Status(fiber.StatusOK).SendString("success")
	} else {
		return c.Status(fiber.StatusUnauthorized).SendString("invalid Request")
	}
}
