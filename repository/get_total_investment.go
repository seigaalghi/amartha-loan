package repository

func (repo *PostgresLoanRepository) GetTotalInvestedAmount(loanID string) (float64, error) {
	var total float64
	result := repo.db.Raw("SELECT COALESCE(SUM(amount), 0) FROM investments WHERE loan_id = ?", loanID).Scan(&total)
	if result.Error != nil {
		return 0, result.Error
	}
	return total, nil
}
