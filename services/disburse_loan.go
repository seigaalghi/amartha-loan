package services

import (
	"net/http"
	"time"

	"github.com/seigaalghi/amartha-loan/common/errs"
	"github.com/seigaalghi/amartha-loan/models"
)

func (s *LoanService) DisburseLoan(loanID string, disbursement models.Disbursement) error {
	loan, err := s.repo.GetLoanByID(loanID)
	if err != nil {
		return err
	}
	if loan.State != models.Invested {
		return errs.NewCustomError(http.StatusBadRequest, "loan can only be disbursed from invested state")
	}
	loan.State = models.Disbursed
	loan.UpdatedAt = time.Now()
	err = s.repo.UpdateLoan(loan)
	if err != nil {
		return errs.NewGeneralError(err)
	}
	return s.repo.SaveDisbursement(disbursement)
}
