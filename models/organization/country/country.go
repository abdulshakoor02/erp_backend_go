package country

import (
	"gorm.io/gorm"
	"time"
)

type Country struct {
	ID             string         `json:"id" gorm:"type:string;size:100;primary_key;default:gen_random_uuid()"`
	Name           string         `json:"name" gorm:"unique;size:300"`
	Code           string         `json:"code" gorm:"unique;size:100"`
	Currency       string         `json:"currency" gorm:"unique;size:100"`
	CurrencyName   string         `json:"currency_name" gorm:"unique;size:100"`
	CurrencySymbol string         `json:"currency_symbol" gorm:"unique;size:50"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
