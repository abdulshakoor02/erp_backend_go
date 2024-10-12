package employees

import (
	"time"

	"github.com/abdul/erp_backend/models/organization/branch"
	"github.com/abdul/erp_backend/models/organization/country"
	"github.com/abdul/erp_backend/models/organization/role"
	tenant "github.com/abdul/erp_backend/models/organization/tenants"
	"gorm.io/gorm"
)

type Employees struct {
	ID         string `json:"id"          gorm:"type:string;size:100;primary_key;default:gen_random_uuid()"`
	FirstName  string `json:"first_name"  gorm:"type:string;size:100"`
	LastName   string `json:"last_name"   gorm:"type:string;size:100"`
	Phone      string `json:"phone"       gorm:"type:string;size:100;unique"`
	Email      string `json:"email"       gorm:"type:string;size:100;unique"`
	Password   string `json:"password"    gorm:"type:string;size:100"`
	RoleId     string `json:"role_id"`
	Role       role.Role
	BranchId   string `json:"branch_id"`
	Branch     branch.Branch
	CountryId  string `json:"country_id"`
	Country    country.Country
	CreatedBy  string `json:"created_by"  gorm:"type:string;size:100"`
	ModifiedBy string `json:"modified_by" gorm:"type:string;size:100"`
	TenantId   string `json:"tenant_id"`
	Tenant     tenant.Tenant
	Status     string         `json:"status"      gorm:"type:string;size:100"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"  gorm:"index"`
}
