package reciepts

import (
	"time"

	"github.com/abdul/erp_backend/models/organization/invoice"
	"gorm.io/gorm"
)

type Reciepts struct {
	ID         string          `json:"id"          gorm:"type:string;size:100;primary_key;default:gen_random_uuid()"`
	RecieptNo  int64           `json:"reciept_no"  gorm:"type:int;autoIncrement;index:idx_reciept"`
	InvoiceId  string          `json:"invoice_id"  gorm:"type:string"`
	Invoice    invoice.Invoice `json:"invoice"`
	AmountPaid float64         `json:"amount_paid" gorm:"type:float"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
	DeletedAt  gorm.DeletedAt  `json:"deleted_at"  gorm:"index"`
}
