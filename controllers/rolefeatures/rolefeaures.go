package rolefeaturesController

import (
	"time"

	"github.com/abdul/erp_backend/controllers/genericHandler"
	"github.com/abdul/erp_backend/models/organization/rolefeatures"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RoleFeatures struct {
	ID         string  `json:"id"`
	RoleId     string  `json:"role_id"`
	Role       Role    `json:"role"`
	TenantId   string  `json:"tenant_id"`
	FeatureId  string  `json:"feature_id"`
	Feature    Feature `json:"feature"`
	CreatedBy  string
	ModifiedBy string
	DeletedAt  gorm.DeletedAt
}

type Role struct {
	ID   string `json:"role_id"`
	Name string `json:"role_name"`
}
type Feature struct {
	ID        string    `json:"feature_id"`
	Name      string    `json:"feature_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt
}

func CreateRolefeatures(c *fiber.Ctx) error {
	return genericHandler.CreateHandler[rolefeatures.RoleFeatures](c)
}

func FindRolefeatures(c *fiber.Ctx) error {

	return genericHandler.FindHandler[RoleFeatures](c)
}

func FindRolefeaturesAssociated(c *fiber.Ctx) error {

	return genericHandler.FindAssociatedHandler[RoleFeatures](c)
}

func UpdateRolefeatures(c *fiber.Ctx) error {

	return genericHandler.UpdateHandler[RoleFeatures, rolefeatures.RoleFeatures](c)
}

func DeleteRolefeatures(c *fiber.Ctx) error {

	return genericHandler.DeleteHandler[RoleFeatures, rolefeatures.RoleFeatures](c)
}
