package rolefeatures

import (
	"time"

	"github.com/abdul/erp_backend/models/organization/features"
	"github.com/abdul/erp_backend/models/organization/role"
	"github.com/abdul/erp_backend/models/organization/tenants"
	"gorm.io/gorm"
)

type RoleFeatures struct {
	ID         string `json:"id" gorm:"type:string;size:100;primary_key;default:gen_random_uuid()"`
	RoleId     string `json:"role_id"`
	Role       role.Role
	TenantId   string `json:"tenant_id"`
	Tenant     tenant.Tenant
	FeatureId  string `json:"feature_id"`
	Feature    features.Features
	CreatedBy  string         `json:"created_by" gorm:"type:string;size:100"`
	ModifiedBy string         `json:"modified_by" gorm:"type:string;size:100"`
	Status     string         `json:"status" gorm:"type:string;size:100"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
