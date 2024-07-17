package controllers

import "github.com/gin-gonic/gin"

type LoanControllerInterface interface {
	ApproveLoan(ctx *gin.Context)
	InvestInLoan(ctx *gin.Context)
	DisburseLoan(ctx *gin.Context)
	CreateLoan(ctx *gin.Context)
}
