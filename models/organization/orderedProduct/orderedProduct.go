package orderedProduct

import (
	"time"

	"github.com/abdul/erp_backend/models/organization/invoice"
	"github.com/abdul/erp_backend/models/organization/products"
	"gorm.io/gorm"
)

type OrderedProduct struct {
	ID        string            `json:"id"         gorm:"type:string;size:100;primary_key;default:gen_random_uuid()"`
	InvoiceId string            `json:"invoice_id" gorm:"type:string"`
	Invoice   invoice.Invoice   `json:"invoice"`
	ProductId string            `json:"product_id" gorm:"type:string"`
	Product   products.Products `json:"product"`
	Price     float64           `json:"price"      gorm:"type:float"`
	Quantity  int64             `json:"quantity"   gorm:"type:int"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	DeletedAt gorm.DeletedAt    `json:"deleted_at" gorm:"index"`
}
