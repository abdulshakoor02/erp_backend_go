package designationController

import (
	"github.com/abdul/erp_backend/controllers/genericHandler"
	"github.com/abdul/erp_backend/models/organization/designation"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Designation struct {
	ID         string
	Name       string
	CreatedBy  string
	ModifiedBy string
	TenantId   string
	DeletedAt  gorm.DeletedAt
}

func CreateDesignation(c *fiber.Ctx) error {
	return genericHandler.CreateHandler[designation.Designation](c)
}

func FindDesignation(c *fiber.Ctx) error {

	return genericHandler.FindHandler[Designation](c)
}

func UpdateDesignation(c *fiber.Ctx) error {

	return genericHandler.UpdateHandler[Designation, designation.Designation](c)
}

func DeleteDesignation(c *fiber.Ctx) error {

	return genericHandler.DeleteHandler[Designation, designation.Designation](c)
}
