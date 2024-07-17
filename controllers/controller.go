// controllers/loan_controller.go
package controllers

import (
	"github.com/seigaalghi/amartha-loan/services"
)

type LoanController struct {
	service services.LoanServiceInterface
}

func NewLoanController(service services.LoanServiceInterface) LoanControllerInterface {
	return &LoanController{service: service}
}
