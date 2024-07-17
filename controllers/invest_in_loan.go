package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seigaalghi/amartha-loan/common/errs"
	"github.com/seigaalghi/amartha-loan/common/helpers"
	"github.com/seigaalghi/amartha-loan/models"
)

func (c *LoanController) InvestInLoan(ctx *gin.Context) {
	var investment models.Investment
	if err := ctx.ShouldBindJSON(&investment); err != nil {
		helpers.GinErrorHandler(ctx, errs.NewCustomErrorWithCause(http.StatusBadRequest, err.Error(), err))
		return
	}
	loanID := ctx.Param("loanID")
	err := c.service.InvestInLoan(loanID, investment)
	if err != nil {
		helpers.GinErrorHandler(ctx, err)
		return
	}

	helpers.GinSuccessResponse(ctx, http.StatusOK, "Investment added successfully", nil)
}
