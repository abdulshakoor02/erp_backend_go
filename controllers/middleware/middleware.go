package middleware

import (
	// "encoding/json"
	// "fmt"
	//
	// "github.com/abdul/erp_backend/database/dbAdapter"

	"encoding/json"
	"fmt"

	"github.com/abdul/erp_backend/logger"

	"github.com/abdul/erp_backend/config"
	"github.com/abdul/erp_backend/models/user"
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
			if path == "/auth" {
				response, err := json.Marshal(user.UserData{
					Id:       1,
					Email:    "test@test.com",
					Username: "test",
					Fullname: "test name",
					Role:     "admin",
				})
				if err != nil {
					log.Err(err).Msgf("failed to sign token : %v", err)
					return c.Status(fiber.StatusUnauthorized).SendString("invalid Request")
				}
				return c.Status(fiber.StatusOK).SendString(string(response))
			}
			return c.Next()
		} else {
			return c.Status(fiber.StatusUnauthorized).SendString("invalid Request")
		}
	}
}
