package tenantsController

import (
	"github.com/abdul/erp_backend/controllers/genericHandler"
	tenant "github.com/abdul/erp_backend/models/organization/tenants"
	"github.com/gofiber/fiber/v2"
	// "time"
)

type Tenants struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Website   string `json:"website"`
	CountryId string `json:"country_id"`
	Status    string `json:"status"`
}

func CreateTenants(c *fiber.Ctx) error {
	return genericHandler.CreateHandler[tenant.Tenant](c)
}

func FindTenants(c *fiber.Ctx) error {

	return genericHandler.FindHandler[Tenants](c)
}

func UpdateTenants(c *fiber.Ctx) error {

	return genericHandler.UpdateHandler[Tenants, tenant.Tenant](c)
}

func DeleteTenants(c *fiber.Ctx) error {

	return genericHandler.DeleteHandler[Tenants, tenant.Tenant](c)
}
