// services/loan_service.go
package services

import (
	"github.com/seigaalghi/amartha-loan/repository"
)

type LoanService struct {
	repo repository.LoanRepository
}

func NewLoanService(repo repository.LoanRepository) LoanServiceInterface {
	return &LoanService{repo: repo}
}
