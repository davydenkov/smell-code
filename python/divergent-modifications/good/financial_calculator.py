class FinancialCalculator:
    def calculate_interest(self, principal: float, rate: float, time: int) -> float:
        return principal * rate * time

    def calculate_tax(self, income: float, deductions: float) -> float:
        taxable_income = income - deductions
        if taxable_income <= 50000:
            return taxable_income * 0.1
        elif taxable_income <= 100000:
            return 5000 + (taxable_income - 50000) * 0.2
        else:
            return 15000 + (taxable_income - 100000) * 0.3
