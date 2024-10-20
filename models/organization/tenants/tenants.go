package tenant

import (
	"time"

	"github.com/abdul/erp_backend/models/organization/country"
	"gorm.io/gorm"
)

type Tenant struct {
	ID        string `json:"id"         gorm:"type:string;size:100;primary_key;default:gen_random_uuid()"`
	Name      string `json:"name"       gorm:"type:string;size:100"`
	Phone     string `json:"phone"      gorm:"type:string;size:100;unique"`
	Email     string `json:"email"      gorm:"type:string;size:100;unique"`
	Website   string `json:"website"    gorm:"type:string;size:300;unique"`
	CountryId string `json:"country_id"`
	Country   country.Country
	Status    string         `json:"status"     gorm:"type:string;size:100"`
	Logo      string         `json:"logo"       gorm:"type:string;size:100;unique"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
