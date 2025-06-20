package products

import (
	"time"

	tenant "github.com/abdul/erp_backend/models/organization/tenants"
	"gorm.io/gorm"
)

type Products struct {
	ID        string         `json:"id"         gorm:"type:string;size:100;primary_key;default:gen_random_uuid()"`
	Name      string         `json:"name"       gorm:"type:string;size:100;unique"`
	Price     float64        `json:"price"      gorm:"type:float"`
	Desc      string         `json:"desc"       gorm:"type:text"`
	TenantId  string         `json:"tenant_id"  gorm:"type:string;size:100"`
	Tenant    tenant.Tenant  `json:"tenant"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
