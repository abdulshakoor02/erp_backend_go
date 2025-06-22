package recieptView

import (
	"time"
)

type RecieptView struct {
	ID            string    `json:"id"`
	InvoiceId     string    `json:"invoice_id"`
	AmountPaid    string    `json:"amount_paid"`
	LeadId        string    `json:"lead_id"`
	Total         string    `json:"total"`
	PendingAmount string    `json:"pending_amount"`
	Discount      string    `json:"discount"`
	RecieptNo     int64     `json:"recipet_no"`
	TenantId      string    `json:"tenant_id"        `
	LeadName      string    `json:"lead_name"      `
	LeadEmail     string    `json:"lead_email"        `
	LeadMobile    string    `json:"lead_mobile" `
	BranchName    string    `json:"branch_name" `
	BranchMobile  string    `json:"branch_mobile" `
	BranchAddress string    `json:"branch_address" `
	CreatedAt     time.Time `json:"created_at"`
}
