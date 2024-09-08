package regionController

import (
	"github.com/abdul/erp_backend/controllers/genericHandler"
	"github.com/abdul/erp_backend/models/organization/region"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Region struct {
	ID         string
	Name       string
	CreatedBy  string
	ModifiedBy string
	TenantId   string
	Status     string
	DeletedAt  gorm.DeletedAt
}

func CreateRegion(c *fiber.Ctx) error {
	return genericHandler.CreateHandler[region.Region](c)
}

func FindRegion(c *fiber.Ctx) error {

	return genericHandler.FindHandler[Region](c)
}

func UpdateRegion(c *fiber.Ctx) error {

	return genericHandler.UpdateHandler[Region, region.Region](c)
}

func DeleteRegion(c *fiber.Ctx) error {

	return genericHandler.DeleteHandler[Region, region.Region](c)
}
