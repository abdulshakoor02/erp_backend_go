package middleware

import (
	// "encoding/json"
	// "fmt"
	//
	// "github.com/abdul/erp_backend/database/dbAdapter"

	"encoding/json"

	"github.com/abdul/erp_backend/logger"

	"github.com/abdul/erp_backend/config"
	employeesController "github.com/abdul/erp_backend/controllers/employees"
	rolefeaturesController "github.com/abdul/erp_backend/controllers/rolefeatures"
	"github.com/abdul/erp_backend/database/dbAdapter"
	"github.com/abdul/erp_backend/models/user"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var log = logger.Logger

type Query struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Employees struct {
	employeesController.Employees
	Password string `json:"password"`
}

func AuthHandler(c *fiber.Ctx) error {
	path := c.Path()
	if path == "/login" {
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

		var resp Employees
		var Where Query
		email, ok := parsedClaim["email"].(string)
		if !ok {
			// Handle the case where the assertion fails
			log.Err(err).Msgf("failed to sign token : %v", err)
			return c.Status(fiber.StatusUnauthorized).SendString("invalid Request")
		}
		password, ok := parsedClaim["password"].(string)
		if !ok {
			// Handle the case where the assertion fails
			log.Err(err).Msgf("failed to sign token : %v", err)
			return c.Status(fiber.StatusUnauthorized).SendString("invalid Request")
		}
		Where.Email = email
		Where.Password = password

		if err := dbAdapter.DB.Model(&resp).InnerJoins("Role").Where(Where).Find(&resp).Error; err != nil {
			log.Info().Msgf("error  %v", err)
			return c.Status(fiber.StatusBadRequest).SendString("invalid request")
		}
		c.Locals("tenant_id", resp.TenantId)
		log.Info().Msgf("tenant id : %v for the current request\n", resp.TenantId)

		if resp.Email != "" {
			var featureList []rolefeaturesController.RoleFeatures
			if err := dbAdapter.DB.Model(&featureList).InnerJoins("Role").InnerJoins("Feature").Where("\"Role\".name = ?", resp.Role.Name).Find(&featureList).Error; err != nil {
				log.Info().Msgf("error  %v", err)
				return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
			}
			features := []string{}
			for _, v := range featureList {
				features = append(features, v.Feature.Name)
			}

			if path == "/auth" {
				response, err := json.Marshal(user.UserData{
					Id:       resp.ID,
					Email:    resp.Email,
					Username: resp.Email,
					Fullname: resp.FirstName,
					Role:     resp.Role.Name,
					Features: features,
					TenantId: resp.TenantId,
				})
				if err != nil {
					log.Err(err).Msgf("failed to sign token : %v", err)
					return c.Status(fiber.StatusUnauthorized).SendString("UnAuthorized")
				}
				return c.Status(fiber.StatusOK).SendString(string(response))
			}
			return c.Next()
		} else {
			return c.Status(fiber.StatusBadRequest).SendString("UnAuthorized User! User not found")
		}

		if parsedClaim["email"] == "test@test.com" && parsedClaim["password"] == "test@123" {
			if path == "/auth" {
				response, err := json.Marshal(user.UserData{
					Id:       "",
					Email:    "test@test.com",
					Username: "test",
					Fullname: "test name",
					Role:     "admin",
				})
				if err != nil {
					log.Err(err).Msgf("failed to sign token : %v", err)
					return c.Status(fiber.StatusUnauthorized).SendString("UnAuthorized")
				}
				return c.Status(fiber.StatusOK).SendString(string(response))
			}
			return c.Next()
		} else {
			return c.Status(fiber.StatusUnauthorized).SendString("UnAuthorized")
		}
	}
}
