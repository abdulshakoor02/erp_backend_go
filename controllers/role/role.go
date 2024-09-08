package roleController

import (
	"github.com/abdul/erp_backend/controllers/genericHandler"
	"github.com/abdul/erp_backend/models/organization/role"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Role struct {
	ID         string
	Name       string
	CreatedBy  string
	ModifiedBy string
	Status     string
	DeletedAt  gorm.DeletedAt
}

func CreateRole(c *fiber.Ctx) error {
	return genericHandler.CreateHandler[role.Role](c)
}

func FindRole(c *fiber.Ctx) error {

	return genericHandler.FindHandler[Role](c)
}

func UpdateRole(c *fiber.Ctx) error {

	return genericHandler.UpdateHandler[Role, role.Role](c)
}

func DeleteRole(c *fiber.Ctx) error {

	return genericHandler.DeleteHandler[Role, role.Role](c)
}
