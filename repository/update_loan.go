package repository

import "github.com/seigaalghi/amartha-loan/models"

func (repo *PostgresLoanRepository) UpdateLoan(loan *models.Loan) error {
	result := repo.db.Exec("UPDATE loans SET borrower_id = ?, principal_amount = ?, rate = ?, roi = ?, agreement_link = ?, state = ?, updated_at = ? WHERE id = ?",
		loan.BorrowerID, loan.PrincipalAmount, loan.Rate, loan.ROI, loan.AgreementLink, loan.State, loan.UpdatedAt, loan.ID)
	return result.Error
}
