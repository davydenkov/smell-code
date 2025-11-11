package main

import (
	"fmt"
	"time"
)

type FinancialCalculator interface {
	CalculateInterest(principal, rate float64, time int) float64
	CalculateTax(income, deductions float64) float64
}

type TransactionRepository interface {
	SaveTransaction(userId int, amount float64, transactionType string) (int64, error)
	GetUserBalance(userId int) float64
	GetMonthlyTransactions(userId, month, year int) []map[string]interface{}
	GetYearlyIncome(userId, year int) float64
	GetYearlyDeductions(userId, year int) float64
}

type UserRepository interface {
	UpdateUserProfile(userId int, name, email string) error
	GetUserEmail(userId int) (string, error)
}

type ReportGenerator interface {
	GenerateMonthlyReport(userId, month, year int) []map[string]interface{}
	GenerateTaxReport(userId, year int) map[string]interface{}
}

type EmailService interface {
	SendMonthlyStatement(email string, balance float64, monthlyReport []map[string]interface{})
}

type FinancialService struct {
	calculator           FinancialCalculator
	transactionRepository TransactionRepository
	userRepository       UserRepository
	reportGenerator      ReportGenerator
	emailService         EmailService
}

func NewFinancialService(
	calculator FinancialCalculator,
	transactionRepo TransactionRepository,
	userRepo UserRepository,
	reportGen ReportGenerator,
	emailSvc EmailService,
) *FinancialService {
	return &FinancialService{
		calculator:            calculator,
		transactionRepository: transactionRepo,
		userRepository:       userRepo,
		reportGenerator:      reportGen,
		emailService:         emailSvc,
	}
}

// Financial calculations - delegated to FinancialCalculator
func (fs FinancialService) CalculateInterest(principal, rate float64, time int) float64 {
	return fs.calculator.CalculateInterest(principal, rate, time)
}

func (fs FinancialService) CalculateTax(income, deductions float64) float64 {
	return fs.calculator.CalculateTax(income, deductions)
}

// Database operations - delegated to repositories
func (fs FinancialService) SaveTransaction(userId int, amount float64, transactionType string) (int64, error) {
	return fs.transactionRepository.SaveTransaction(userId, amount, transactionType)
}

func (fs FinancialService) GetUserBalance(userId int) float64 {
	return fs.transactionRepository.GetUserBalance(userId)
}

func (fs FinancialService) UpdateUserProfile(userId int, name, email string) error {
	return fs.userRepository.UpdateUserProfile(userId, name, email)
}

// Reporting - delegated to ReportGenerator
func (fs FinancialService) GenerateMonthlyReport(userId, month, year int) []map[string]interface{} {
	return fs.reportGenerator.GenerateMonthlyReport(userId, month, year)
}

func (fs FinancialService) GenerateTaxReport(userId, year int) map[string]interface{} {
	return fs.reportGenerator.GenerateTaxReport(userId, year)
}

// Email notifications - delegated to EmailService
func (fs FinancialService) SendMonthlyStatement(userId int) error {
	email, err := fs.userRepository.GetUserEmail(userId)
	if err != nil {
		return err
	}

	balance := fs.GetUserBalance(userId)
	currentTime := time.Now()
	report := fs.GenerateMonthlyReport(userId, int(currentTime.Month()), currentTime.Year())

	fs.emailService.SendMonthlyStatement(email, balance, report)
	return nil
}
