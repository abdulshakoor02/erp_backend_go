package leadCategory

import (
	"time"

	tenant "github.com/abdul/erp_backend/models/organization/tenants"
	"gorm.io/gorm"
)

type LeadCategory struct {
	ID        string         `json:"id"         gorm:"type:string;size:100;primary_key;default:gen_random_uuid()"`
	Name      string         `json:"name"       gorm:"type:string;size:100;unique"`
	TenantId  string         `json:"tenant_id"  gorm:"type:string;size:100"`
	Tenant    tenant.Tenant  `json:"tenant"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
