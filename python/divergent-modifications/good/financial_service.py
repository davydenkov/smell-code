from datetime import datetime

class FinancialService:
    def __init__(self, db):
        self.calculator = FinancialCalculator()
        self.transaction_repository = TransactionRepository(db)
        self.user_repository = UserRepository(db)
        self.report_generator = ReportGenerator(self.transaction_repository, self.calculator)
        self.email_service = EmailService()

    # Financial calculations - delegated to FinancialCalculator
    def calculate_interest(self, principal: float, rate: float, time: int) -> float:
        return self.calculator.calculate_interest(principal, rate, time)

    def calculate_tax(self, income: float, deductions: float) -> float:
        return self.calculator.calculate_tax(income, deductions)

    # Database operations - delegated to repositories
    def save_transaction(self, user_id: int, amount: float, transaction_type: str) -> int:
        return self.transaction_repository.save_transaction(user_id, amount, transaction_type)

    def get_user_balance(self, user_id: int) -> float:
        return self.transaction_repository.get_user_balance(user_id)

    def update_user_profile(self, user_id: int, name: str, email: str) -> None:
        self.user_repository.update_user_profile(user_id, name, email)

    # Reporting - delegated to ReportGenerator
    def generate_monthly_report(self, user_id: int, month: int, year: int):
        return self.report_generator.generate_monthly_report(user_id, month, year)

    def generate_tax_report(self, user_id: int, year: int):
        return self.report_generator.generate_tax_report(user_id, year)

    # Email notifications - delegated to EmailService
    def send_monthly_statement(self, user_id: int) -> None:
        email = self.user_repository.get_user_email(user_id)
        if not email:
            return

        balance = self.get_user_balance(user_id)
        current_month = datetime.now().month
        current_year = datetime.now().year
        report = self.generate_monthly_report(user_id, current_month, current_year)

        self.email_service.send_monthly_statement(email, balance, report)
