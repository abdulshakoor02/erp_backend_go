package features

import (
	"gorm.io/gorm"
	"time"
)

type Features struct {
	ID         string         `json:"id" gorm:"type:string;size:100;primary_key;default:gen_random_uuid()"`
	Name       string         `json:"name" gorm:"type:string;size:100;unique"`
	CreatedBy  string         `json:"created_by" gorm:"type:string;size:100"`
	ModifiedBy string         `json:"modified_by" gorm:"type:string;size:100"`
	Status     string         `json:"status" gorm:"type:string;size:100"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
