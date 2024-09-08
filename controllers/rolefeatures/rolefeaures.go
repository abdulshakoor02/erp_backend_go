package rolefeaturesController

import (
	"github.com/abdul/erp_backend/controllers/genericHandler"
	"github.com/abdul/erp_backend/models/organization/rolefeatures"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Rolefeatures struct {
	ID         string
	RoleId     string
	TenantId   string
	FeatureId  string
	CreatedBy  string
	ModifiedBy string
	Status     string
	DeletedAt  gorm.DeletedAt
}

func CreateRolefeatures(c *fiber.Ctx) error {
	return genericHandler.CreateHandler[rolefeatures.RoleFeatures](c)
}

func FindRolefeatures(c *fiber.Ctx) error {

	return genericHandler.FindHandler[Rolefeatures](c)
}

func UpdateRolefeatures(c *fiber.Ctx) error {

	return genericHandler.UpdateHandler[Rolefeatures, rolefeatures.RoleFeatures](c)
}

func DeleteRolefeatures(c *fiber.Ctx) error {

	return genericHandler.DeleteHandler[Rolefeatures, rolefeatures.RoleFeatures](c)
}
