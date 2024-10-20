package addtionalInfoController

import (
	"github.com/abdul/erp_backend/controllers/genericHandler"
	"github.com/abdul/erp_backend/models/organization/additionalInfo"
	"github.com/gofiber/fiber/v2"
)

type AdditionalInfo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"price"`
	LeadId      string `json:"lead_id"`
}

func Create(c *fiber.Ctx) error {
	return genericHandler.CreateHandler[additionalInfo.AdditionalInfo](c)
}

func Find(c *fiber.Ctx) error {

	return genericHandler.FindHandler[additionalInfo.AdditionalInfo](c)
}

func Update(c *fiber.Ctx) error {

	return genericHandler.UpdateHandler[AdditionalInfo, additionalInfo.AdditionalInfo](c)
}

func Delete(c *fiber.Ctx) error {

	return genericHandler.DeleteHandler[AdditionalInfo, additionalInfo.AdditionalInfo](c)
}
