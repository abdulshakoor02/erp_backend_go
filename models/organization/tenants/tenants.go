package tenant

import (
	"github.com/abdul/erp_backend/models/organization/country"
	"gorm.io/gorm"
	"time"
)

type Tenant struct {
	ID         string `json:"id" gorm:"type:string;size:100;primary_key;default:gen_random_uuid()"`
	Name       string `json:"name" gorm:"type:string;size:100"`
	Phone      string `json:"phone" gorm:"type:string;size:100;unique"`
	Email      string `json:"email" gorm:"type:string;size:100;unique"`
	Website    string `json:"website" gorm:"type:string;size:300;unique"`
	CountryId  string `json:"country_id"`
	Country    country.Country
	CreatedBy  string         `json:"created_by" gorm:"type:string;size:100"`
	ModifiedBy string         `json:"modified_by" gorm:"type:string;size:100"`
	Status     string         `json:"status" gorm:"type:string;size:100"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
