package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/seigaalghi/amartha-loan/common/errs"
	"github.com/seigaalghi/amartha-loan/models"
)

func (s *LoanService) CreateLoan(loan *models.Loan) error {
	loan.ID = uuid.New().String()
	loan.State = models.Proposed
	loan.CreatedAt = time.Now()
	loan.UpdatedAt = time.Now()

	err := s.repo.CreateLoan(loan)
	if err != nil {
		return errs.NewGeneralError(err)
	}

	return nil
}
