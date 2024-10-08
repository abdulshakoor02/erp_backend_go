package tenantsController

import (
	"github.com/abdul/erp_backend/controllers/genericHandler"
	tenant "github.com/abdul/erp_backend/models/organization/tenants"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	// "time"
)

type Tenants struct {
	gorm.Model
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Website   string    `json:"website"`
	CountryId string    `json:"country_id"`
	Country   Countries `json:"country"`
	Status    string    `json:"status"`
}

type Countries struct {
	gorm.Model
	ID   string `json:"country_id"`
	Name string `json:"country_name"`
}

func CreateTenants(c *fiber.Ctx) error {
	return genericHandler.CreateHandler[tenant.Tenant](c)
}

func FindTenants(c *fiber.Ctx) error {

	return genericHandler.FindHandler[Tenants](c)
}

func FindAssociatedTenants(c *fiber.Ctx) error {

	return genericHandler.FindAssociatedHandler[Tenants](c)
}

func UpdateTenants(c *fiber.Ctx) error {

	return genericHandler.UpdateHandler[Tenants, tenant.Tenant](c)
}

func DeleteTenants(c *fiber.Ctx) error {

	return genericHandler.DeleteHandler[Tenants, tenant.Tenant](c)
}
