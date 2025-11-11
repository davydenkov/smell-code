class EmailService:
    def send_monthly_statement(self, email: str, balance: float, report) -> None:
        subject = 'Monthly Financial Statement'
        message = f"Your current balance: ${balance:.2f}\n\n"
        message += "Transactions this month:\n"

        for row in report:
            transaction_type, total, count = row
            message += f"- {transaction_type}: {count} transactions, total: ${total:.2f}\n"

        # Simulate sending email
        print(f"Sending email to {email}")
        print(f"Subject: {subject}")
        print(f"Message: {message}")
