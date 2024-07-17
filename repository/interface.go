package repository

import "github.com/seigaalghi/amartha-loan/models"

type LoanRepository interface {
	GetLoanByID(id string) (*models.Loan, error)
	UpdateLoan(loan *models.Loan) error
	SaveApproval(approval models.Approval) error
	SaveInvestment(investment models.Investment) error
	GetTotalInvestedAmount(loanID string) (float64, error)
	SaveDisbursement(disbursement models.Disbursement) error
	CreateLoan(loan *models.Loan) error
}
