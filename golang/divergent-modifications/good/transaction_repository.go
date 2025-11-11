package main

import (
	"database/sql"
	"log"
)

type transactionRepository struct {
	db Database
}

func NewTransactionRepository(db Database) TransactionRepository {
	return &transactionRepository{db: db}
}

func (tr transactionRepository) SaveTransaction(userId int, amount float64, transactionType string) (int64, error) {
	result, err := tr.db.Exec(
		"INSERT INTO transactions (user_id, amount, type, created_at) VALUES (?, ?, ?, NOW())",
		userId, amount, transactionType,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (tr transactionRepository) GetUserBalance(userId int) float64 {
	rows, err := tr.db.Query("SELECT SUM(amount) as balance FROM transactions WHERE user_id = ?", userId)
	if err != nil {
		log.Printf("Error querying balance: %v", err)
		return 0
	}
	defer rows.Close()

	if rows.Next() {
		var balance *float64
		rows.Scan(&balance)
		if balance != nil {
			return *balance
		}
	}
	return 0
}

func (tr transactionRepository) GetMonthlyTransactions(userId, month, year int) []map[string]interface{} {
	rows, err := tr.db.Query(`
		SELECT type, SUM(amount) as total, COUNT(*) as count
		FROM transactions
		WHERE user_id = ? AND MONTH(created_at) = ? AND YEAR(created_at) = ?
		GROUP BY type
	`, userId, month, year)

	if err != nil {
		log.Printf("Error getting monthly transactions: %v", err)
		return nil
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		var transactionType string
		var total float64
		var count int
		rows.Scan(&transactionType, &total, &count)
		results = append(results, map[string]interface{}{
			"type":  transactionType,
			"total": total,
			"count": count,
		})
	}
	return results
}

func (tr transactionRepository) GetYearlyIncome(userId, year int) float64 {
	rows, err := tr.db.Query(`
		SELECT SUM(amount) as total_income
		FROM transactions
		WHERE user_id = ? AND type = 'income' AND YEAR(created_at) = ?
	`, userId, year)

	if err != nil {
		log.Printf("Error getting yearly income: %v", err)
		return 0
	}
	defer rows.Close()

	if rows.Next() {
		var income *float64
		rows.Scan(&income)
		if income != nil {
			return *income
		}
	}
	return 0
}

func (tr transactionRepository) GetYearlyDeductions(userId, year int) float64 {
	rows, err := tr.db.Query(`
		SELECT SUM(amount) as total_deductions
		FROM transactions
		WHERE user_id = ? AND type = 'deduction' AND YEAR(created_at) = ?
	`, userId, year)

	if err != nil {
		log.Printf("Error getting yearly deductions: %v", err)
		return 0
	}
	defer rows.Close()

	if rows.Next() {
		var deductions *float64
		rows.Scan(&deductions)
		if deductions != nil {
			return *deductions
		}
	}
	return 0
}
