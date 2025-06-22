package orderedProductsView

import (
	"time"
)

type OrderedProductsView struct {
	ID           string    `json:"id"`
	InvoiceId    string    `json:"invoice_id"`
	Quantity     string    `json:"quantity"`
	ProductId    string    `json:"product_id"`
	ProductDesc  string    `json:"product_desc"`
	ProductPrice string    `json:"product_price"`
	ProductName  string    `json:"product_name"`
	CreatedAt    time.Time `json:"created_at"`
}
