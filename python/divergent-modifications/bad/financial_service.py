import sqlite3
from datetime import datetime

class FinancialService:
    def __init__(self, db):
        self.db = db

    # Financial calculations - changes when business rules change
    def calculate_interest(self, principal, rate, time):
        return principal * rate * time

    def calculate_tax(self, income, deductions):
        taxable_income = income - deductions
        if taxable_income <= 50000:
            return taxable_income * 0.1
        elif taxable_income <= 100000:
            return 5000 + (taxable_income - 50000) * 0.2
        else:
            return 15000 + (taxable_income - 100000) * 0.3

    # Database operations - changes when data schema changes
    def save_transaction(self, user_id, amount, transaction_type):
        cursor = self.db.cursor()
        cursor.execute(
            "INSERT INTO transactions (user_id, amount, type, created_at) VALUES (?, ?, ?, ?)",
            (user_id, amount, transaction_type, datetime.now())
        )
        self.db.commit()
        return cursor.lastrowid

    def get_user_balance(self, user_id):
        cursor = self.db.cursor()
        cursor.execute("SELECT SUM(amount) as balance FROM transactions WHERE user_id = ?", (user_id,))
        result = cursor.fetchone()
        return result[0] if result[0] is not None else 0

    def update_user_profile(self, user_id, name, email):
        cursor = self.db.cursor()
        cursor.execute("UPDATE users SET name = ?, email = ? WHERE id = ?", (name, email, user_id))
        self.db.commit()

    # Reporting - changes when reporting requirements change
    def generate_monthly_report(self, user_id, month, year):
        cursor = self.db.cursor()
        cursor.execute("""
            SELECT type, SUM(amount) as total, COUNT(*) as count
            FROM transactions
            WHERE user_id = ? AND strftime('%m', created_at) = ? AND strftime('%Y', created_at) = ?
            GROUP BY type
        """, (user_id, f"{month:02d}", str(year)))
        return cursor.fetchall()

    def generate_tax_report(self, user_id, year):
        cursor = self.db.cursor()

        cursor.execute("""
            SELECT SUM(amount) as total_income
            FROM transactions
            WHERE user_id = ? AND type = 'income' AND strftime('%Y', created_at) = ?
        """, (user_id, str(year)))
        income = cursor.fetchone()[0] or 0

        cursor.execute("""
            SELECT SUM(amount) as total_deductions
            FROM transactions
            WHERE user_id = ? AND type = 'deduction' AND strftime('%Y', created_at) = ?
        """, (user_id, str(year)))
        deductions = cursor.fetchone()[0] or 0

        tax_owed = self.calculate_tax(income, deductions)

        return {
            'year': year,
            'total_income': income,
            'total_deductions': deductions,
            'tax_owed': tax_owed
        }

    # Email notifications - changes when communication requirements change
    def send_monthly_statement(self, user_id, email):
        balance = self.get_user_balance(user_id)
        current_month = datetime.now().month
        current_year = datetime.now().year
        report = self.generate_monthly_report(user_id, current_month, current_year)

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
