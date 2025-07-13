package invoiceController

import (
	"encoding/json"
	"sync"

	"github.com/abdul/erp_backend/controllers/genericHandler"
	"github.com/abdul/erp_backend/database/dbAdapter"
	"github.com/abdul/erp_backend/logger"
	"github.com/abdul/erp_backend/models/organization/invoice"
	"github.com/abdul/erp_backend/models/organization/orderProductsView"
	"github.com/abdul/erp_backend/models/organization/orderedProduct"
	"github.com/abdul/erp_backend/models/organization/reciepts"
	"github.com/abdul/erp_backend/models/organization/recieptsView"
	"github.com/gofiber/fiber/v2"
)

type CreateInvoice struct {
	LeadId     string                          `json:"lead_id"`
	Total      float64                         `json:"total"`
	AmountPaid float64                         `json:"amount_paid"`
	Discount   float64                         `json:"discount"`
	TenantId   string                          `json:"tenant_id"`
	Products   []orderedProduct.OrderedProduct `json:"products"`
}

type GetOneInvoice struct {
	InvoiceId string `json:"invoice_id"`
	RecieptId string `json:"reciept_id"`
}

type OneInvoiceResponse struct {
	Reciept recieptView.RecieptView                   `json:"reciept"`
	Orders  []orderedProductsView.OrderedProductsView `json:"orders"`
}

type CreateInvoiceResponse struct {
	Invoice invoice.Invoice   `json:"invoice"`
	Reciept reciepts.Reciepts `json:"reciept"`
}

var log = logger.Logger

func Create(c *fiber.Ctx) error {
	db := dbAdapter.DB
	var Payload CreateInvoice
	err := json.Unmarshal(c.Body(), &Payload)
	if err != nil {
		log.Info().Msgf("error  %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
	}
	tenantId := c.Locals("tenant_id")

	var Invoice invoice.Invoice
	tenant_id, ok := tenantId.(string)
	if ok && tenant_id != "" {
		Invoice.TenantId = tenant_id
	} else {
		return c.Status(fiber.StatusBadRequest).SendString("cannot create invoice without a tenant")
	}

	Invoice.AmountPaid = Payload.AmountPaid
	Invoice.Total = Payload.Total
	Invoice.PendingAmount = Payload.Total - Payload.AmountPaid - Payload.Discount
	Invoice.LeadId = Payload.LeadId
	Invoice.Discount = Payload.Discount

	if err := db.Create(&Invoice).Error; err != nil {
		log.Info().Msgf("failed to create invoice  %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("failed to create invoice")
	}

	var Products []orderedProduct.OrderedProduct

	for _, prods := range Payload.Products {
		prods.InvoiceId = Invoice.ID
		Products = append(Products, prods)
	}

	if err := db.Create(&Products).Error; err != nil {
		log.Info().Msgf("failed to create ordered products  %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("failed to create products")
	}

	var Reciepts reciepts.Reciepts
	Reciepts.InvoiceId = Invoice.ID
	Reciepts.AmountPaid = Payload.AmountPaid

	if err := db.Create(&Reciepts).Error; err != nil {
		log.Info().Msgf("failed to create reciepts  %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("failed to create reciepts")
	}

	var response CreateInvoiceResponse
	response.Invoice = Invoice
	response.Reciept = Reciepts

	newJSONData2, err := json.Marshal(response)
	if err != nil {
		log.Info().Msgf("error  %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
	}

	return c.Status(fiber.StatusOK).SendString(string(newJSONData2))
}

func FindOne(c *fiber.Ctx) error {
	db := dbAdapter.DB
	var Payload GetOneInvoice
	err := json.Unmarshal(c.Body(), &Payload)
	if err != nil {
		log.Info().Msgf("error  %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
	}
	tenantId := c.Locals("tenant_id")

	tenant_id, ok := tenantId.(string)
	if ok && tenant_id != "" {
	} else {
		return c.Status(fiber.StatusBadRequest).SendString("cannot create invoice without a tenant")
	}

	var wg sync.WaitGroup
	wg.Add(2)

	var (
		reciepts    []recieptView.RecieptView
		orderdProds []orderedProductsView.OrderedProductsView

		recieptsErr, ordrdProdsErr error
	)

	go func() {
		defer wg.Done()
		recieptsErr = db.Table("reciepts_view").Where("id = ?", Payload.RecieptId).Find(&reciepts).Error
		log.Info().Msgf("fetched reciepts")
	}()

	go func() {
		defer wg.Done()
		ordrdProdsErr = db.Table("ordered_products_view").Where("invoice_id = ?", Payload.InvoiceId).Find(&orderdProds).Error
		log.Info().Msgf("fetched orderdProds")
	}()

	wg.Wait()
	// Check errors after all goroutines finish
	if recieptsErr != nil {
		log.Info().Msgf("failed to create invoice  %v", recieptsErr)
		return c.Status(fiber.StatusBadRequest).SendString("cannot create invoice without a tenant")
	}
	if ordrdProdsErr != nil {
		log.Info().Msgf("failed to create invoice  %v", ordrdProdsErr)
		return c.Status(fiber.StatusBadRequest).SendString("cannot create invoice without a tenant")
	}

	var response OneInvoiceResponse
	response.Reciept = reciepts[0]
	response.Orders = orderdProds

	newJSONData2, err := json.Marshal(response)
	if err != nil {
		log.Info().Msgf("error  %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
	}

	return c.Status(fiber.StatusOK).SendString(string(newJSONData2))
}

func FindReciepts(c *fiber.Ctx) error {

	return genericHandler.FindHandler[recieptView.RecieptView](c)
}
