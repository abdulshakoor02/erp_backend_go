package invoiceController

import (
	"encoding/json"
	"fmt"
	"math"
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

type CreateReciept struct {
	InvoiceId  string  `json:"invoice_id"`
	AmountPaid float64 `json:"amount_paid"`
}

type GetOneInvoice struct {
	InvoiceId string `json:"invoice_id"`
	RecieptId string `json:"reciept_id"`
}

type OneInvoiceResponse struct {
	Reciept     recieptView.RecieptView                   `json:"reciept"`
	RecieptList []reciepts.Reciepts                       `json:"reciept_list"`
	Orders      []orderedProductsView.OrderedProductsView `json:"orders"`
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
	pendingAmount := Payload.Total - (Payload.AmountPaid + Payload.Discount)
	pendingAmount = math.Max(pendingAmount, 0)
	Invoice.PendingAmount = pendingAmount
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
	wg.Add(3)

	var (
		recpts      []recieptView.RecieptView
		recieptList []reciepts.Reciepts
		orderdProds []orderedProductsView.OrderedProductsView

		recieptsErr, ordrdProdsErr, recieptsListErr error
	)

	go func() {
		defer wg.Done()
		db := dbAdapter.DB
		recieptsErr = db.Table("reciepts_view").Where("id = ?", Payload.RecieptId).Find(&recpts).Error
		log.Info().Msgf("fetched reciepts")
	}()

	go func() {
		defer wg.Done()
		db := dbAdapter.DB
		ordrdProdsErr = db.Table("ordered_products_view").Where("invoice_id = ?", Payload.InvoiceId).Find(&orderdProds).Error
		log.Info().Msgf("fetched orderdProds")
	}()

	go func() {
		defer wg.Done()
		db := dbAdapter.DB
		recieptsListErr = db.Where("invoice_id = ?", Payload.InvoiceId).Find(&recieptList).Error
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

	if recieptsListErr != nil {
		log.Info().Msgf("failed to create invoice  %v", recieptsListErr)
		return c.Status(fiber.StatusBadRequest).SendString("cannot create invoice without a tenant")
	}

	var response OneInvoiceResponse
	response.Reciept = recpts[0]
	response.Orders = orderdProds
	response.RecieptList = recieptList

	newJSONData2, err := json.Marshal(response)
	if err != nil {
		log.Info().Msgf("error  %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
	}

	return c.Status(fiber.StatusOK).SendString(string(newJSONData2))
}

func FindReciepts(c *fiber.Ctx) error {
	var where map[string]interface{}

	err := json.Unmarshal(c.Body(), &where)
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
		reciepts []recieptView.RecieptView
		count    int64

		recieptsErr, countErr error
	)

	go func() {
		defer wg.Done()

		db := dbAdapter.DB
		for key, value := range where {
			clause := fmt.Sprintf("\"reciepts_view\".\"%v\" = ? ", key)
			db = db.Where(clause, value)
		}
		countErr = db.Table("reciepts_view").Count(&count).Error
		log.Info().Msgf("fetched reciepts count")
	}()

	go func() {
		defer wg.Done()

		db := dbAdapter.DB
		for key, value := range where {
			clause := fmt.Sprintf("\"reciepts_view\".\"%v\" = ? ", key)
			db = db.Where(clause, value)
		}
		recieptsErr = db.Table("reciepts_view").Find(&reciepts).Error
		log.Info().Msgf("fetched reciepts")
	}()

	wg.Wait()
	// Check errors after all goroutines finish
	if recieptsErr != nil {
		log.Info().Msgf("failed to fetch invoice  %v", recieptsErr)
		return c.Status(fiber.StatusBadRequest).SendString("unable to fetch invoice")
	}

	if countErr != nil {
		log.Info().Msgf("failed to fetch invoice count  %v", countErr)
		return c.Status(fiber.StatusBadRequest).SendString("unable to fetch count")
	}
	var response struct {
		Count int64                     `json:"count"`
		Data  []recieptView.RecieptView `json:"data"`
	}

	response.Count = count
	response.Data = reciepts

	// dbAdapter.DB.Find(&result)
	newJSONData2, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Print the JSON object and authorization header
	// fmt.Println(string(newJSONData))

	// Return a 200 OK response

	return c.Status(fiber.StatusOK).SendString(string(newJSONData2))
}

func FindInvoices(c *fiber.Ctx) error {

	return genericHandler.FindAssociatedHandler[invoice.Invoice](c)
}

func GenerateReciept(c *fiber.Ctx) error {

	db := dbAdapter.DB
	var Payload CreateReciept
	err := json.Unmarshal(c.Body(), &Payload)
	if err != nil {
		log.Info().Msgf("error  %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
	}
	var Invoice invoice.Invoice
	var InvoiceData invoice.Invoice
	db.Where("id = ?", Payload.InvoiceId).First(&Invoice)
	pendingAmount := Invoice.Total - (Invoice.AmountPaid + Invoice.Discount + Payload.AmountPaid)
	log.Info().Msgf("before pending amount %v", pendingAmount)
	pendingAmount = math.Max(pendingAmount, 0)
	log.Info().Msgf("after pending amount %v", pendingAmount)

	InvoiceData.AmountPaid = Payload.AmountPaid + Invoice.AmountPaid
	InvoiceData.PendingAmount = pendingAmount

	var wg sync.WaitGroup
	wg.Add(2)

	var recieptErr, invoiceErr error

	go func() {
		defer wg.Done()
		invoiceErr = db.Where("id = ?", Payload.InvoiceId).Select("amount_paid", "pending_amount").Updates(&InvoiceData).Error
	}()

	var Reciepts reciepts.Reciepts
	Reciepts.InvoiceId = Payload.InvoiceId
	Reciepts.AmountPaid = Payload.AmountPaid

	go func() {
		defer wg.Done()
		recieptErr = db.Create(&Reciepts).Error
	}()

	wg.Wait()

	if recieptErr != nil {
		log.Info().Msgf("failed to create reciepts  %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("failed to create reciepts")
	}

	if invoiceErr != nil {
		log.Info().Msgf("failed to update invoice  %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("failed to update invoice")
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
