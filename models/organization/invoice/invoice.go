package invoice

import (
	"time"

	"github.com/abdul/erp_backend/models/organization/leads"
	tenant "github.com/abdul/erp_backend/models/organization/tenants"
	"gorm.io/gorm"
)

type Invoice struct {
	ID            string         `json:"id"             gorm:"type:string;size:100;primary_key;default:gen_random_uuid()"`
	InvoiceNo     int64          `json:"invoice_no"     gorm:"type:int;autoIncrement;index:idx_invoice"`
	LeadId        string         `json:"lead_id"        gorm:"type:string"`
	Lead          leads.Leads    `json:"leads"`
	Total         float64        `json:"total"          gorm:"type:float"`
	AmountPaid    float64        `json:"amount_paid"    gorm:"type:float"`
	PendingAmount float64        `json:"pending_amount" gorm:"type:float"`
	TenantId      string         `json:"tenant_id"      gorm:"type:string"`
	Tenant        tenant.Tenant  `json:"tenant"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"     gorm:"index"`
}
