package repository

import "github.com/seigaalghi/amartha-loan/models"

func (repo *PostgresLoanRepository) GetLoanByID(id string) (*models.Loan, error) {
	var loan models.Loan
	result := repo.db.Raw("SELECT * FROM loans WHERE id = ?", id).Scan(&loan)
	if result.Error != nil {
		return nil, result.Error
	}
	return &loan, nil
}
