package middleware

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

func AuthHandler(c *fiber.Ctx) error {
	path := c.Path()
	if path == "/login" {
		fmt.Println("unprotected route")
		return c.Next()
	} else {
		// to do handle token verification here
		secretKey := []byte(config.SECRET_KEY)
		tokenValue := c.Get("token", "")
		token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			log.Err(err).Msgf("failed to sign token : %v", err)
			return c.Status(fiber.StatusUnauthorized).SendString("invalid Request")
		}

		parsedClaim := token.Claims.(jwt.MapClaims)
		if parsedClaim["email"] == "test@test.com" && parsedClaim["password"] == "test@123" {
			fmt.Println("user authorised")
			return c.Next()
		} else {
			return c.Status(fiber.StatusUnauthorized).SendString("invalid Request")
		}
	}
}
