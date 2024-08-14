package repository

import "github.com/seigaalghi/amartha-loan/models"

func (repo *PostgresLoanRepository) SaveDisbursement(disbursement models.Disbursement) error {
	result := repo.db.Exec("INSERT INTO disbursements (id, loan_id, field_officer_id, disbursement_date, agreement_letter) VALUES (?, ?, ?, ?, ?)",
		disbursement.ID, disbursement.LoanID, disbursement.FieldOfficerID, disbursement.DisbursementDate, disbursement.AgreementLetter)
	return result.Error
}
