package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seigaalghi/amartha-loan/common/errs"
	"github.com/seigaalghi/amartha-loan/common/helpers"
	"github.com/seigaalghi/amartha-loan/models"
)

func (c *LoanController) DisburseLoan(ctx *gin.Context) {
	var disbursement models.Disbursement
	if err := ctx.ShouldBindJSON(&disbursement); err != nil {
		helpers.GinErrorHandler(ctx, errs.NewCustomErrorWithCause(http.StatusBadRequest, err.Error(), err))
		return
	}
	loanID := ctx.Param("loanID")
	err := c.service.DisburseLoan(loanID, disbursement)
	if err != nil {
		helpers.GinErrorHandler(ctx, err)
		return
	}

	helpers.GinSuccessResponse(ctx, http.StatusOK, "Loan disbursed successfully", nil)
}
