package leadView

import (
	"time"

	"gorm.io/gorm"
)

type LeadView struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	Email          string         `json:"email"`
	Mobile         string         `json:"mobile"`
	Address        string         `json:"address"`
	CountryId      string         `json:"country_id"`
	TenantId       string         `json:"tenant_id"        `
	EmployeeId     string         `json:"employee_id"      `
	BranchId       string         `json:"branch_id"        `
	LeadCategoryId string         `json:"lead_category_id" `
	ProductId      string         `json:"product_id"`
	EmployeeName   string         `json:"employee_name"`
	BranchName     string         `json:"branch_name"`
	Client         bool           `json:"client"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"       `
}
