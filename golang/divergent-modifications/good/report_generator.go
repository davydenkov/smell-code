package main

type reportGenerator struct {
	transactionRepository TransactionRepository
	calculator           FinancialCalculator
}

func NewReportGenerator(transactionRepo TransactionRepository, calculator FinancialCalculator) ReportGenerator {
	return &reportGenerator{
		transactionRepository: transactionRepo,
		calculator:           calculator,
	}
}

func (rg reportGenerator) GenerateMonthlyReport(userId, month, year int) []map[string]interface{} {
	return rg.transactionRepository.GetMonthlyTransactions(userId, month, year)
}

func (rg reportGenerator) GenerateTaxReport(userId, year int) map[string]interface{} {
	income := rg.transactionRepository.GetYearlyIncome(userId, year)
	deductions := rg.transactionRepository.GetYearlyDeductions(userId, year)
	taxOwed := rg.calculator.CalculateTax(income, deductions)

	return map[string]interface{}{
		"year":             year,
		"total_income":     income,
		"total_deductions": deductions,
		"tax_owed":         taxOwed,
	}
}
