package leadCategoryController

import (
	"github.com/abdul/erp_backend/controllers/genericHandler"
	"github.com/abdul/erp_backend/models/organization/leadCategory"
	"github.com/gofiber/fiber/v2"
)

type LeadCategory struct {
	ID       string `json:"id"        gorm:"type:string;size:100"`
	TenantId string `json:"tenant_id"`
	Name     string `json:"name"      gorm:"type:string;size:100"`
}

func Create(c *fiber.Ctx) error {
	return genericHandler.CreateHandler[leadCategory.LeadCategory](c)
}

func Find(c *fiber.Ctx) error {

	return genericHandler.FindHandler[leadCategory.LeadCategory](c)
}

func Update(c *fiber.Ctx) error {

	return genericHandler.UpdateHandler[LeadCategory, leadCategory.LeadCategory](c)
}

func Delete(c *fiber.Ctx) error {

	return genericHandler.DeleteHandler[LeadCategory, leadCategory.LeadCategory](c)
}
