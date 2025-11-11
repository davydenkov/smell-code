package main

type financialCalculator struct{}

func NewFinancialCalculator() FinancialCalculator {
	return &financialCalculator{}
}

func (fc financialCalculator) CalculateInterest(principal, rate float64, time int) float64 {
	return principal * rate * float64(time)
}

func (fc financialCalculator) CalculateTax(income, deductions float64) float64 {
	taxableIncome := income - deductions
	if taxableIncome <= 50000 {
		return taxableIncome * 0.1
	} else if taxableIncome <= 100000 {
		return 5000 + (taxableIncome-50000)*0.2
	} else {
		return 15000 + (taxableIncome-100000)*0.3
	}
}
