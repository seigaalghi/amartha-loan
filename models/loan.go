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
	ID              string    `json:"id"`
	BorrowerID      string    `json:"borrower_id"`
	PrincipalAmount float64   `json:"principal_amount"`
	Rate            float64   `json:"rate"`
	ROI             float64   `json:"roi"`
	AgreementLink   string    `json:"agreement_link"`
	State           LoanState `json:"state"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Approval struct {
	LoanID       string    `json:"loan_id"`
	ValidatorID  string    `json:"validator_id"`
	ApprovalDate time.Time `json:"approval_date"`
	ProofPicture string    `json:"proof_picture"`
}

type Investment struct {
	LoanID     string    `json:"loan_id"`
	InvestorID string    `json:"investor_id"`
	Amount     float64   `json:"amount"`
	InvestedAt time.Time `json:"invested_at"`
}

type Disbursement struct {
	LoanID           string    `json:"loan_id"`
	FieldOfficerID   string    `json:"field_officer_id"`
	DisbursementDate time.Time `json:"disbursement_date"`
	AgreementLetter  string    `json:"agreement_letter"`
}
