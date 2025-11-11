package main

import (
	"database/sql"
	"fmt"
)

// Mock database interface for demonstration
type Database interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

func main() {
	// This would normally use a real database connection
	var db Database // mock database

	// Create all the separate services
	calculator := NewFinancialCalculator()
	transactionRepo := NewTransactionRepository(db)
	userRepo := NewUserRepository(db)
	reportGen := NewReportGenerator(transactionRepo, calculator)
	emailSvc := NewEmailService()

	// Create the main financial service with all dependencies
	fs := NewFinancialService(calculator, transactionRepo, userRepo, reportGen, emailSvc)

	// Example usage
	interest := fs.CalculateInterest(1000, 0.05, 2)
	fmt.Printf("Interest: $%.2f\n", interest)

	tax := fs.CalculateTax(75000, 10000)
	fmt.Printf("Tax: $%.2f\n", tax)

	fmt.Println("Each responsibility is now handled by a separate, focused component:")
	fmt.Println("- FinancialCalculator: handles business calculations")
	fmt.Println("- TransactionRepository: handles transaction data operations")
	fmt.Println("- UserRepository: handles user data operations")
	fmt.Println("- ReportGenerator: handles report generation")
	fmt.Println("- EmailService: handles email communications")
}
