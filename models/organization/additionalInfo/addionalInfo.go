package additionalInfo

import (
	"time"

	"github.com/abdul/erp_backend/models/organization/leads"
	"gorm.io/gorm"
)

type AdditionalInfo struct {
	ID          string         `json:"id"          gorm:"type:string;size:100;primary_key;default:gen_random_uuid()"`
	Title       string         `json:"title"       gorm:"type:string;size:100;index:idx_title"`
	Description string         `json:"description" gorm:"type:text"`
	LeadId      string         `json:"lead_id"     gorm:"type:string"`
	Lead        leads.Leads    `json:"leads"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"  gorm:"index"`
}
