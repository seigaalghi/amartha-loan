package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seigaalghi/amartha-loan/common/configs"
	"github.com/seigaalghi/amartha-loan/common/errs"
	"go.uber.org/zap"
)

func GinErrorHandler(ctx *gin.Context, err error) {
	status := http.StatusInternalServerError
	message := err.Error()

	localErr, ok := err.(*errs.CustomError)
	if ok {
		status = localErr.Code
		message = localErr.Message
		err = localErr.Err
	}

	requestID, _ := ctx.Get("RequestID")

	configs.Logger.Info("error", zap.String("request_id", requestID.(string)), zap.Error(err))

	ctx.JSON(status, gin.H{"message": message})
}
