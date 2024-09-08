package employeesController

import (
	"github.com/abdul/erp_backend/controllers/genericHandler"
	"github.com/abdul/erp_backend/models/organization/employees"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Employees struct {
	ID         string
	FirstName  string
	LastName   string
	Phone      string
	Email      string
	Password   string
	Country    string
	CreatedBy  string
	ModifiedBy string
	TenantId   string
	Status     string
	DeletedAt  gorm.DeletedAt
}

func CreateEmployees(c *fiber.Ctx) error {
	return genericHandler.CreateHandler[employees.Employees](c)
}

func FindEmployees(c *fiber.Ctx) error {

	return genericHandler.FindHandler[Employees](c)
}

func UpdateEmployees(c *fiber.Ctx) error {

	return genericHandler.UpdateHandler[Employees, employees.Employees](c)
}

func DeleteEmployees(c *fiber.Ctx) error {

	return genericHandler.DeleteHandler[Employees, employees.Employees](c)
}
