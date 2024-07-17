package repository

import "github.com/seigaalghi/amartha-loan/models"

func (repo *PostgresLoanRepository) CreateLoan(loan *models.Loan) error {
	result := repo.db.Exec("INSERT INTO loans (id, borrower_id, principal_amount, rate, roi, agreement_link, state, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		loan.ID, loan.BorrowerID, loan.PrincipalAmount, loan.Rate, loan.ROI, loan.AgreementLink, loan.State, loan.CreatedAt, loan.UpdatedAt)
	return result.Error
}
