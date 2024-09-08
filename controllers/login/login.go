package login

import (
	"encoding/json"
	"fmt"

	"github.com/abdul/erp_backend/config"
	"github.com/abdul/erp_backend/logger"
	"github.com/abdul/erp_backend/models/user"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var log = logger.Logger

func Login(c *fiber.Ctx) error {
	secretKey := []byte(config.SECRET_KEY)
	u := new(user.User)
	if err := c.BodyParser(u); err != nil {
		return err
	}
	if u.Email == "test@test.com" && u.Password == "test@123" {
		claims := jwt.MapClaims{
			"email":    u.Email,
			"password": u.Password,
		}
		jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedToken, err := jwtToken.SignedString(secretKey)
		if err != nil {
			log.Err(err).Msgf("failed to sign token : %v", err)
		}
		fmt.Println(signedToken)

		// Return a 200 OK response
		response, err2 := json.Marshal(struct {
			Token string `json:"token"`
		}{Token: signedToken})
		if err2 != nil {
			log.Err(err2).Msgf("failed to sign token : %v", err2)
			return c.Status(fiber.StatusUnauthorized).SendString("invalid Request")
		}
		return c.Status(fiber.StatusOK).SendString(string(response))
	} else {
		return c.Status(fiber.StatusUnauthorized).SendString("invalid credentials")
	}

}
