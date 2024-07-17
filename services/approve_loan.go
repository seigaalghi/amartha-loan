package services

import (
	"net/http"
	"time"

	"github.com/seigaalghi/amartha-loan/common/errs"
	"github.com/seigaalghi/amartha-loan/models"
)

func (s *LoanService) ApproveLoan(loanID string, approval models.Approval) error {
	loan, err := s.repo.GetLoanByID(loanID)
	if err != nil {
		return err
	}
	if loan.State != models.Proposed {
		return errs.NewCustomError(http.StatusBadRequest, "loan can only be approved from proposed state")
	}
	loan.State = models.Approved
	loan.UpdatedAt = time.Now()
	err = s.repo.UpdateLoan(loan)
	if err != nil {
		return errs.NewGeneralError(err)
	}
	return s.repo.SaveApproval(approval)
}
