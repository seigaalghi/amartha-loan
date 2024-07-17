// repository/postgres_loan_repository.go
package repository

import (
	"gorm.io/gorm"
)

type PostgresLoanRepository struct {
	db *gorm.DB
}

func NewPostgresLoanRepository(db *gorm.DB) *PostgresLoanRepository {
	return &PostgresLoanRepository{db: db}
}
