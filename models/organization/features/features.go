package features

import (
	"time"

	"gorm.io/gorm"
)

type Features struct {
	ID        string         `json:"id"         gorm:"type:string;size:100;primary_key;default:gen_random_uuid()"`
	Name      string         `json:"name"       gorm:"type:string;size:100;unique"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
