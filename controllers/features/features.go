package featuresController

import (
	"github.com/abdul/erp_backend/controllers/genericHandler"
	"github.com/abdul/erp_backend/models/organization/features"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Features struct {
	ID         string
	Name       string
	CreatedBy  string
	ModifiedBy string
	Status     string
	DeletedAt  gorm.DeletedAt
}

func CreateFeatures(c *fiber.Ctx) error {
	return genericHandler.CreateHandler[features.Features](c)
}

func FindFeatures(c *fiber.Ctx) error {

	return genericHandler.FindHandler[Features](c)
}

func UpdateFeatures(c *fiber.Ctx) error {

	return genericHandler.UpdateHandler[Features, features.Features](c)
}

func DeleteFeatures(c *fiber.Ctx) error {

	return genericHandler.DeleteHandler[Features, features.Features](c)
}
