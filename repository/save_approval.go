package repository

import "github.com/seigaalghi/amartha-loan/models"

func (repo *PostgresLoanRepository) SaveApproval(approval models.Approval) error {
	result := repo.db.Exec("INSERT INTO approvals (loan_id, validator_id, approval_date, proof_picture) VALUES (?, ?, ?, ?)",
		approval.LoanID, approval.ValidatorID, approval.ApprovalDate, approval.ProofPicture)
	return result.Error
}
