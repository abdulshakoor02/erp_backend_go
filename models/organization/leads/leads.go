package leads

import (
	"time"

	"github.com/abdul/erp_backend/models/organization/branch"
	"github.com/abdul/erp_backend/models/organization/country"
	"github.com/abdul/erp_backend/models/organization/employees"
	"github.com/abdul/erp_backend/models/organization/leadCategory"
	tenant "github.com/abdul/erp_backend/models/organization/tenants"
	"gorm.io/gorm"
)

type Leads struct {
	ID             string          `json:"id"               gorm:"type:string;size:100;primary_key;default:gen_random_uuid()"`
	Name           string          `json:"name"             gorm:"type:string;size:100"`
	Email          string          `json:"email"            gorm:"type:string;size:100;index:,unique,composite:tenant_email"`
	Mobile         string          `json:"mobile"           gorm:"type:string;size:100;index:,unique,composite:tenant_mobile"`
	Address        string          `json:"address"          gorm:"type:text"`
	CountryId      string          `json:"country_id"       gorm:"type:string"`
	Country        country.Country `json:"country"`
	TenantId       string          `json:"tenant_id"        gorm:"type:string;size:100;index:,unique,composite:tenant_email;index:,unique,composite:tenant_mobile"`
	Tenant         tenant.Tenant
	EmployeeId     string                    `json:"employee_id"      gorm:"type:string;size:100"`
	Employee       employees.Employees       `json:"employee"`
	BranchId       string                    `json:"branch_id"        gorm:"type:string;size:100"`
	Branch         branch.Branch             `json:"branch"`
	LeadCategoryId string                    `json:"lead_category_id" gorm:"type:string;size:100"`
	LeadCategory   leadCategory.LeadCategory `json:"lead_category"`
	ProductId      string                    `json:"product_id"       gorm:"type:string;size:100"`
	Client         bool                      `json:"client"           gorm:"type:boolean;index:idx_client"`
	CreatedAt      time.Time                 `json:"created_at"`
	UpdatedAt      time.Time                 `json:"updated_at"`
	DeletedAt      gorm.DeletedAt            `json:"deleted_at"       gorm:"index"`
}
