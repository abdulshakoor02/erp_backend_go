package appointmentsController

import (
	"github.com/abdul/erp_backend/controllers/genericHandler"
	"github.com/abdul/erp_backend/models/organization/appointments"
	"github.com/gofiber/fiber/v2"
)

type Appointments struct {
	ID       string `json:"id"        gorm:"type:string;size:100;primary_key;default:gen_random_uuid()"`
	DateTime string `json:"date_time" gorm:"type:string;size:100"`
	LeadId   string `json:"lead_id"   gorm:"type:string"`
}

func Create(c *fiber.Ctx) error {
	return genericHandler.CreateHandler[appointments.Appointments](c)
}

func Find(c *fiber.Ctx) error {

	return genericHandler.FindHandler[appointments.Appointments](c)
}

func Update(c *fiber.Ctx) error {

	return genericHandler.UpdateHandler[Appointments, appointments.Appointments](c)
}

func Delete(c *fiber.Ctx) error {

	return genericHandler.DeleteHandler[Appointments, appointments.Appointments](c)
}
