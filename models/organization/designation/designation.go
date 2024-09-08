package designation

import (
	"github.com/abdul/erp_backend/models/organization/tenants"
	"gorm.io/gorm"
	"time"
)

type Designation struct {
	ID         string `json:"id" gorm:"type:string;size:100;primary_key;default:gen_random_uuid()"`
	Name       string `json:"name" gorm:"type:string;size:200;unique"`
	CreatedBy  string `json:"created_by" gorm:"type:string;size:100"`
	ModifiedBy string `json:"modified_by" gorm:"type:string;size:100"`
	TenantId   string `json:"tenant_id"`
	Tenant     tenant.Tenant
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
