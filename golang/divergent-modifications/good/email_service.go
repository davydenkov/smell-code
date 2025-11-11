package main

import "fmt"

type emailService struct{}

func NewEmailService() EmailService {
	return &emailService{}
}

func (es emailService) SendMonthlyStatement(email string, balance float64, monthlyReport []map[string]interface{}) {
	subject := "Monthly Financial Statement"
	message := fmt.Sprintf("Your current balance: $%.2f\n\n", balance)
	message += "Transactions this month:\n"

	for _, row := range monthlyReport {
		message += fmt.Sprintf("- %s: %d transactions, total: $%.2f\n",
			row["type"], row["count"], row["total"])
	}

	fmt.Printf("Sending email to %s:\nSubject: %s\n%s\n", email, subject, message)
	// In real implementation, would send actual email
}
