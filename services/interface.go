package services

import "github.com/seigaalghi/amartha-loan/models"

type LoanServiceInterface interface {
	ApproveLoan(loanID string, approval models.Approval) error
	InvestInLoan(loanID string, investment models.Investment) error
	DisburseLoan(loanID string, disbursement models.Disbursement) error
	CreateLoan(loan *models.Loan) error
}
