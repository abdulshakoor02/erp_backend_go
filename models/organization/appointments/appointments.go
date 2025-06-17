package appointments

import (
	"time"

	"github.com/abdul/erp_backend/models/organization/leads"
	"gorm.io/gorm"
)

type Appointments struct {
	ID        string         `json:"id"         gorm:"type:string;size:100;primary_key;default:gen_random_uuid()"`
	DateTime  string         `json:"date_time"  gorm:"type:string;size:100"`
	LeadId    string         `json:"lead_id"    gorm:"type:string"`
	Lead      leads.Leads    `json:"leads"`
	Status    string         `json:"status" gorm:"type:string"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
