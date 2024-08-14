// models/loan.go
package models

import (
	"time"
)

type LoanState string

const (
	Proposed  LoanState = "proposed"
	Approved  LoanState = "approved"
	Invested  LoanState = "invested"
	Disbursed LoanState = "disbursed"
)

type Loan struct {
	ID              string    `json:"id" gorm:"primaryKey"`
	BorrowerID      string    `json:"borrower_id" binding:"required"`
	PrincipalAmount float64   `json:"principal_amount" binding:"required"`
	Rate            float64   `json:"rate" binding:"required"`
	ROI             float64   `json:"roi"`
	AgreementLink   string    `json:"agreement_link"`
	State           LoanState `json:"state"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Approval struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	LoanID       string    `json:"loan_id"`
	ValidatorID  string    `json:"validator_id" binding:"required"`
	ApprovalDate time.Time `json:"approval_date" binding:"required"`
	ProofPicture string    `json:"proof_picture" binding:"required"`
}

type Investment struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	LoanID     string    `json:"loan_id"`
	InvestorID string    `json:"investor_id" binding:"required"`
	Amount     float64   `json:"amount" binding:"required"`
	InvestedAt time.Time `json:"invested_at"`
	ROI        float64   `json:"roi"`
}

type Disbursement struct {
	ID               string    `json:"id" gorm:"primaryKey"`
	LoanID           string    `json:"loan_id"`
	FieldOfficerID   string    `json:"field_officer_id"`
	DisbursementDate time.Time `json:"disbursement_date"`
	AgreementLetter  string    `json:"agreement_letter"`
}
