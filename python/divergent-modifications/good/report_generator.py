class ReportGenerator:
    def __init__(self, transaction_repository, calculator):
        self.transaction_repository = transaction_repository
        self.calculator = calculator

    def generate_monthly_report(self, user_id: int, month: int, year: int):
        return self.transaction_repository.get_monthly_transactions(user_id, month, year)

    def generate_tax_report(self, user_id: int, year: int):
        income = self.transaction_repository.get_yearly_income(user_id, year)
        deductions = self.transaction_repository.get_yearly_deductions(user_id, year)
        tax_owed = self.calculator.calculate_tax(income, deductions)

        return {
            'year': year,
            'total_income': income,
            'total_deductions': deductions,
            'tax_owed': tax_owed
        }
