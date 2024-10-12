package employeesController

import (
	"github.com/abdul/erp_backend/controllers/genericHandler"
	"github.com/abdul/erp_backend/models/organization/employees"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Employees struct {
	gorm.Model
	ID         string    `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	CreatedBy  string    `json:"created_by"`
	ModifiedBy string    `json:"modified_by"`
	CountryId  string    `json:"country_id"`
	RoleId     string    `json:"role_id"`
	BranchId   string    `json:"branch_id"`
	Country    Countries `json:"country"`
	Role       Role      `json:"role"`
	Branch     Branch    `json:"branch"`
	TenantId   string    `json:"tenant_id"`
	Status     string    `json:"status"`
	DeletedAt  gorm.DeletedAt
}

type Countries struct {
	gorm.Model
	ID   string `json:"country_id"`
	Name string `json:"country_name"`
}

type Role struct {
	gorm.Model
	ID   string `json:"role_id"`
	Name string `json:"role_name"`
}

type Branch struct {
	gorm.Model
	ID   string `json:"branch_id"`
	Name string `json:"branch_name"`
}

func CreateEmployees(c *fiber.Ctx) error {
	return genericHandler.CreateHandler[employees.Employees](c)
}

func FindEmployees(c *fiber.Ctx) error {

	return genericHandler.FindHandler[Employees](c)
}

func FindEmployeesAssociated(c *fiber.Ctx) error {

	return genericHandler.FindAssociatedHandler[Employees](c)
}

func UpdateEmployees(c *fiber.Ctx) error {

	return genericHandler.UpdateHandler[Employees, employees.Employees](c)
}

func DeleteEmployees(c *fiber.Ctx) error {

	return genericHandler.DeleteHandler[Employees, employees.Employees](c)
}
