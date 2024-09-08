package branch

import (
	"github.com/abdul/erp_backend/models/organization/country"
	"github.com/abdul/erp_backend/models/organization/region"
	"gorm.io/gorm"
	"time"
)

type Branch struct {
	ID        string `json:"id" gorm:"type:string;size:100;primary_key;default:gen_random_uuid()"`
	Name      string `json:"name" gorm:"type:string;size:100" `
	Mobile    string `json:"mobile" gorm:"type:string;size:100;unique"`
	Email     string `json:"email" gorm:"type:string;size:300;unique"`
	RegionId  string `json:"region_id"`
	Region    region.Region
	Website   string `json:"website" gorm:"type:string;size:100"`
	CountryId string `json:"country_id"`
	Country   country.Country
	// Country    country.Country `json:"country_id" gorm:"column:country_id;:type:string;size:100;foreignKey:country_id;references:id"`
	CreatedBy  string         `json:"createdBy" gorm:"type:string;size:100"`
	ModifiedBy string         `json:"modifiedBy" gorm:"type:string;size:100"`
	Status     string         `json:"status" gorm:"type:string;size:100"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
