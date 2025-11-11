package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

// Mock database interface for demonstration
type Database interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type FinancialService struct {
	db Database
}

func NewFinancialService(db Database) *FinancialService {
	return &FinancialService{db: db}
}

// Financial calculations - changes when business rules change
func (fs FinancialService) CalculateInterest(principal, rate float64, time int) float64 {
	return principal * rate * float64(time)
}

func (fs FinancialService) CalculateTax(income, deductions float64) float64 {
	taxableIncome := income - deductions
	if taxableIncome <= 50000 {
		return taxableIncome * 0.1
	} else if taxableIncome <= 100000 {
		return 5000 + (taxableIncome-50000)*0.2
	} else {
		return 15000 + (taxableIncome-100000)*0.3
	}
}

// Database operations - changes when data schema changes
func (fs FinancialService) SaveTransaction(userId int, amount float64, transactionType string) (int64, error) {
	result, err := fs.db.Exec(
		"INSERT INTO transactions (user_id, amount, type, created_at) VALUES (?, ?, ?, NOW())",
		userId, amount, transactionType,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (fs FinancialService) GetUserBalance(userId int) float64 {
	rows, err := fs.db.Query("SELECT SUM(amount) as balance FROM transactions WHERE user_id = ?", userId)
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

func (fs FinancialService) UpdateUserProfile(userId int, name, email string) error {
	_, err := fs.db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", name, email, userId)
	return err
}

// Reporting - changes when reporting requirements change
func (fs FinancialService) GenerateMonthlyReport(userId, month, year int) []map[string]interface{} {
	rows, err := fs.db.Query(`
		SELECT type, SUM(amount) as total, COUNT(*) as count
		FROM transactions
		WHERE user_id = ? AND MONTH(created_at) = ? AND YEAR(created_at) = ?
		GROUP BY type
	`, userId, month, year)

	if err != nil {
		log.Printf("Error generating monthly report: %v", err)
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

func (fs FinancialService) GenerateTaxReport(userId, year int) map[string]interface{} {
	// Get income
	incomeRows, err := fs.db.Query(`
		SELECT SUM(amount) as total_income
		FROM transactions
		WHERE user_id = ? AND type = 'income' AND YEAR(created_at) = ?
	`, userId, year)

	var income float64
	if err == nil && incomeRows.Next() {
		incomeRows.Scan(&income)
		incomeRows.Close()
	}

	// Get deductions
	deductionRows, err := fs.db.Query(`
		SELECT SUM(amount) as total_deductions
		FROM transactions
		WHERE user_id = ? AND type = 'deduction' AND YEAR(created_at) = ?
	`, userId, year)

	var deductions float64
	if err == nil && deductionRows.Next() {
		deductionRows.Scan(&deductions)
		deductionRows.Close()
	}

	taxOwed := fs.CalculateTax(income, deductions)

	return map[string]interface{}{
		"year":             year,
		"total_income":     income,
		"total_deductions": deductions,
		"tax_owed":         taxOwed,
	}
}

// Email notifications - changes when communication requirements change
func (fs FinancialService) SendMonthlyStatement(userId int, email string) {
	balance := fs.GetUserBalance(userId)
	currentTime := time.Now()
	report := fs.GenerateMonthlyReport(userId, int(currentTime.Month()), currentTime.Year())

	subject := "Monthly Financial Statement"
	message := fmt.Sprintf("Your current balance: $%.2f\n\n", balance)
	message += "Transactions this month:\n"

	for _, row := range report {
		message += fmt.Sprintf("- %s: %d transactions, total: $%.2f\n",
			row["type"], row["count"], row["total"])
	}

	fmt.Printf("Sending email to %s:\nSubject: %s\n%s\n", email, subject, message)
	// In real implementation, would send actual email
}

func main() {
	// This would normally use a real database connection
	var db Database // mock database

	fs := NewFinancialService(db)

	// Example usage
	interest := fs.CalculateInterest(1000, 0.05, 2)
	fmt.Printf("Interest: $%.2f\n", interest)

	tax := fs.CalculateTax(75000, 10000)
	fmt.Printf("Tax: $%.2f\n", tax)
}
