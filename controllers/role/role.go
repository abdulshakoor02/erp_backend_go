package roleController

import (
	"github.com/abdul/erp_backend/controllers/genericHandler"
	"github.com/abdul/erp_backend/models/organization/role"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Role struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CreatedBy  string
	ModifiedBy string
	Status     string `json:"status"`
	DeletedAt  gorm.DeletedAt
}

func CreateRole(c *fiber.Ctx) error {
	return genericHandler.CreateHandler[role.Role](c)
}

func FindRole(c *fiber.Ctx) error {

	return genericHandler.FindAssociatedHandler[role.Role](c)
}

func UpdateRole(c *fiber.Ctx) error {

	return genericHandler.UpdateHandler[Role, role.Role](c)
}

func DeleteRole(c *fiber.Ctx) error {

	return genericHandler.DeleteHandler[Role, role.Role](c)
}
