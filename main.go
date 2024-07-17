// main.go
package main

import (
	"fmt"
	"log"

	"github.com/seigaalghi/amartha-loan/common/configs"
	"github.com/seigaalghi/amartha-loan/common/middlewares"
	"github.com/seigaalghi/amartha-loan/controllers"
	"github.com/seigaalghi/amartha-loan/repository"
	"github.com/seigaalghi/amartha-loan/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	configs.InitLogger()

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable TimeZone=Asia/Jakarta",
		configs.AppConfig.DBUser,
		configs.AppConfig.DBPassword,
		configs.AppConfig.DBName,
		configs.AppConfig.DBPort,
		configs.AppConfig.DBHost,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	loanRepo := repository.NewPostgresLoanRepository(db)
	loanService := services.NewLoanService(loanRepo)
	loanController := controllers.NewLoanController(loanService)

	r := gin.Default()

	r.Use(middlewares.RequestIDMiddleware(configs.Logger))
	r.Use(middlewares.LoggerMiddleware(configs.Logger))

	api := r.Group("/api/v1")
	{
		api.POST("/loans", loanController.CreateLoan)
		api.POST("/loans/:loanID/approve", loanController.ApproveLoan)
		api.POST("/loans/:loanID/invest", loanController.InvestInLoan)
		api.POST("/loans/:loanID/disburse", loanController.DisburseLoan)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
