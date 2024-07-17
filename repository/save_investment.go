package repository

import "github.com/seigaalghi/amartha-loan/models"

func (repo *PostgresLoanRepository) SaveInvestment(investment models.Investment) error {
	result := repo.db.Exec("INSERT INTO investments (loan_id, investor_id, amount, invested_at) VALUES (?, ?, ?, ?)",
		investment.LoanID, investment.InvestorID, investment.Amount, investment.InvestedAt)
	return result.Error
}
