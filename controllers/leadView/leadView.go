package leadViewController

import (
	"github.com/abdul/erp_backend/controllers/genericHandler"
	"github.com/abdul/erp_backend/models/organization/leadView"
	"github.com/gofiber/fiber/v2"
)

func Find(c *fiber.Ctx) error {

	return genericHandler.FindHandler[leadView.LeadView](c)
}
