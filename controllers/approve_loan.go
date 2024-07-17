package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seigaalghi/amartha-loan/common/errs"
	"github.com/seigaalghi/amartha-loan/common/helpers"
	"github.com/seigaalghi/amartha-loan/models"
)

func (c *LoanController) ApproveLoan(ctx *gin.Context) {
	var approval models.Approval
	if err := ctx.ShouldBindJSON(&approval); err != nil {
		helpers.GinErrorHandler(ctx, errs.NewCustomErrorWithCause(http.StatusBadRequest, err.Error(), err))
		return
	}
	loanID := ctx.Param("loanID")
	err := c.service.ApproveLoan(loanID, approval)
	if err != nil {
		helpers.GinErrorHandler(ctx, err)
		return
	}

	helpers.GinSuccessResponse(ctx, http.StatusOK, "Loan approved successfully", nil)
}
