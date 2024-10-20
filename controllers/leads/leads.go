package leadsController

import (
	"github.com/abdul/erp_backend/controllers/genericHandler"
	"github.com/abdul/erp_backend/models/organization/leads"
	"github.com/gofiber/fiber/v2"
)

type Leads struct {
	ID     string `json:"id"     gorm:"type:string;size:100"`
	Name   string `json:"name"   gorm:"type:string;size:100"`
	Email  string `json:"email"  gorm:"type:string;size:100"`
	Mobile string `json:"mobile" gorm:"type:string;size:100"`
}

func Create(c *fiber.Ctx) error {
	return genericHandler.CreateHandler[leads.Leads](c)
}

func Find(c *fiber.Ctx) error {

	return genericHandler.FindAssociatedHandler[leads.Leads](c)
}

func Update(c *fiber.Ctx) error {

	return genericHandler.UpdateHandler[Leads, leads.Leads](c)
}

func Delete(c *fiber.Ctx) error {

	return genericHandler.DeleteHandler[Leads, leads.Leads](c)
}
