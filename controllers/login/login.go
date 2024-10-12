package login

import (
	"encoding/json"

	"github.com/abdul/erp_backend/config"
	employeesController "github.com/abdul/erp_backend/controllers/employees"
	rolefeaturesController "github.com/abdul/erp_backend/controllers/rolefeatures"
	"github.com/abdul/erp_backend/database/dbAdapter"
	"github.com/abdul/erp_backend/logger"
	"github.com/abdul/erp_backend/models/user"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Query struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Employees struct {
	employeesController.Employees
	Password string `json:"password"`
}

var log = logger.Logger

func Login(c *fiber.Ctx) error {
	secretKey := []byte(config.SECRET_KEY)
	u := new(user.User)
	if err := c.BodyParser(u); err != nil {
		return err
	}

	var resp Employees
	var Where Query
	Where.Email = u.Email
	Where.Password = u.Password

	if err := dbAdapter.DB.Model(&resp).InnerJoins("Role").Where(Where).Find(&resp).Error; err != nil {
		log.Info().Msgf("error  %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
	}

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
		claims := jwt.MapClaims{
			"email":    resp.Email,
			"password": resp.Password,
		}

		jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedToken, err := jwtToken.SignedString(secretKey)
		if err != nil {
			log.Err(err).Msgf("failed to sign token : %v", err)
		}

		// Return a 200 OK response
		response, err2 := json.Marshal(struct {
			Token    string        `json:"token"`
			UserData user.UserData `json:"userData"`
		}{
			Token: signedToken,
			UserData: user.UserData{
				Id:       resp.ID,
				Email:    resp.Email,
				Username: resp.Email,
				Fullname: resp.FirstName,
				Role:     resp.Role.Name,
				Features: features,
				TenantId: resp.TenantId,
			},
		})
		if err2 != nil {
			log.Err(err2).Msgf("failed to sign token : %v", err2)
			return c.Status(fiber.StatusUnauthorized).SendString("invalid Request")
		}
		return c.Status(fiber.StatusOK).SendString(string(response))
	} else {
		return c.Status(fiber.StatusUnauthorized).SendString("invalid credentials")
	}

	// if u.Email == "test@test.com" && u.Password == "test@123" {
	// 	claims := jwt.MapClaims{
	// 		"email":    u.Email,
	// 		"password": u.Password,
	// 	}
	// 	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 	signedToken, err := jwtToken.SignedString(secretKey)
	// 	if err != nil {
	// 		log.Err(err).Msgf("failed to sign token : %v", err)
	// 	}
	//
	// 	// Return a 200 OK response
	// 	response, err2 := json.Marshal(struct {
	// 		Token    string        `json:"token"`
	// 		UserData user.UserData `json:"userData"`
	// 	}{
	// 		Token: signedToken,
	// 		UserData: user.UserData{
	// 			Id:       1,
	// 			Email:    "test@test.com",
	// 			Username: "test",
	// 			Fullname: "test name",
	// 			Role:     "admin",
	// 		},
	// 	})
	// 	if err2 != nil {
	// 		log.Err(err2).Msgf("failed to sign token : %v", err2)
	// 		return c.Status(fiber.StatusUnauthorized).SendString("invalid Request")
	// 	}
	// 	return c.Status(fiber.StatusOK).SendString(string(response))
	// } else {
	// 	return c.Status(fiber.StatusUnauthorized).SendString("invalid credentials")
	// }

}
