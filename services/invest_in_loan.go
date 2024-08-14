package services

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/seigaalghi/amartha-loan/common/errs"
	"github.com/seigaalghi/amartha-loan/models"
)

func (s *LoanService) InvestInLoan(loanID string, investment models.Investment) error {
	loan, err := s.repo.GetLoanByID(loanID)
	if err != nil {
		return err
	}
	if loan.State != models.Approved {
		return errs.NewCustomError(http.StatusBadRequest, "loan can only be invested in approved state")
	}
	totalInvested, err := s.repo.GetTotalInvestedAmount(loanID)
	if err != nil {
		return err
	}
	if totalInvested+investment.Amount > loan.PrincipalAmount {
		return errs.NewCustomError(http.StatusBadRequest, "investment amount exceeds loan principal amount")
	}

	roi := calculateROI(investment.Amount, loan.Rate)
	investment.ROI = roi
	investment.LoanID = loanID
	investment.ID = uuid.New().String()
	err = s.repo.SaveInvestment(investment)
	if err != nil {
		return errs.NewGeneralError(err)
	}

	totalInvested += investment.Amount
	if totalInvested == loan.PrincipalAmount {
		loan.State = models.Invested
		loan.UpdatedAt = time.Now()
		err = s.repo.UpdateLoan(loan)
		if err != nil {
			return errs.NewGeneralError(err)
		}
		// upload pdf
		// send email
	}

	return nil
}

func calculateROI(amount, interestRate float64) float64 {
	return amount * (interestRate / 100)
}
