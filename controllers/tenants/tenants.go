package tenantsController

import (
	"github.com/abdul/erp_backend/controllers/genericHandler"
	"github.com/abdul/erp_backend/models/organization/tenants"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	// "time"
)

type TenantsWhere struct {
	ID         string         `json:"id"`
	Name       string         `json:"name"`
	Phone      string         `json:"phone"`
	Email      string         `json:"email"`
	Website    string         `json:"website"`
	CountryId  string         `json:"country_id"`
	CreatedBy  string         `json:"created_by"`
	ModifiedBy string         `json:"modified_by"`
	Status     string         `json:"status"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

func CreateTenants(c *fiber.Ctx) error {
	return genericHandler.CreateHandler[tenant.Tenant](c)
}

func FindTenants(c *fiber.Ctx) error {

	return genericHandler.FindHandler[TenantsWhere](c)
}

func UpdateTenants(c *fiber.Ctx) error {

	return genericHandler.UpdateHandler[TenantsWhere, tenant.Tenant](c)
}

func DeleteTenants(c *fiber.Ctx) error {

	return genericHandler.DeleteHandler[TenantsWhere, tenant.Tenant](c)
}
