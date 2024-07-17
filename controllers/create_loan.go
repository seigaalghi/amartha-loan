package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seigaalghi/amartha-loan/common/errs"
	"github.com/seigaalghi/amartha-loan/common/helpers"
	"github.com/seigaalghi/amartha-loan/models"
)

func (c *LoanController) CreateLoan(ctx *gin.Context) {
	var loan models.Loan
	if err := ctx.ShouldBindJSON(&loan); err != nil {
		helpers.GinErrorHandler(ctx, errs.NewCustomErrorWithCause(http.StatusBadRequest, err.Error(), err))
		return
	}
	err := c.service.CreateLoan(&loan)
	if err != nil {
		helpers.GinErrorHandler(ctx, err)
		return
	}

	helpers.GinSuccessResponse(ctx, http.StatusCreated, "Loan created successfully", loan)
}
