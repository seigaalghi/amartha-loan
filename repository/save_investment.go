package repository

import "github.com/seigaalghi/amartha-loan/models"

func (repo *PostgresLoanRepository) SaveInvestment(investment models.Investment) error {
	result := repo.db.Exec("INSERT INTO investments (id, loan_id, investor_id, amount, invested_at, roi) VALUES (?, ?, ?, ?, ?, ?)",
		investment.ID, investment.LoanID, investment.InvestorID, investment.Amount, investment.InvestedAt, investment.ROI)
	return result.Error
}
