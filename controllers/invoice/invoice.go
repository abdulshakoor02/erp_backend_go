package invoiceController

import (
	"encoding/json"
	"fmt"

	"github.com/abdul/erp_backend/database/dbAdapter"
	"github.com/abdul/erp_backend/logger"
	"github.com/abdul/erp_backend/models/organization/invoice"
	"github.com/abdul/erp_backend/models/organization/orderedProduct"
	"github.com/abdul/erp_backend/models/organization/reciepts"
	"github.com/gofiber/fiber/v2"
)

type CreateInvoice struct {
	LeadId     string                          `json:"lead_id"`
	Total      float64                         `json:"total"`
	AmountPaid float64                         `json:"amount_paid"`
	TenantId   string                          `json:"tenant_id"`
	Products   []orderedProduct.OrderedProduct `json:"products"`
}

var log = logger.Logger

func Create(c *fiber.Ctx) error {
	db := dbAdapter.DB
	var Payload CreateInvoice
	fmt.Print("here")
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
	Invoice.PendingAmount = Payload.Total - Payload.AmountPaid
	Invoice.LeadId = Payload.LeadId

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

	newJSONData2, err := json.Marshal(Invoice)
	if err != nil {
		log.Info().Msgf("error  %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
	}

	return c.Status(fiber.StatusOK).SendString(string(newJSONData2))
}

// func Find(c *fiber.Ctx) error {
//
// 	return genericHandler.FindHandler[additionalInfo.AdditionalInfo](c)
// }
