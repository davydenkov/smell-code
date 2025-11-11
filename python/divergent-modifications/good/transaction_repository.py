from datetime import datetime

class TransactionRepository:
    def __init__(self, db):
        self.db = db

    def save_transaction(self, user_id: int, amount: float, transaction_type: str) -> int:
        cursor = self.db.cursor()
        cursor.execute(
            "INSERT INTO transactions (user_id, amount, type, created_at) VALUES (?, ?, ?, ?)",
            (user_id, amount, transaction_type, datetime.now())
        )
        self.db.commit()
        return cursor.lastrowid

    def get_user_balance(self, user_id: int) -> float:
        cursor = self.db.cursor()
        cursor.execute("SELECT SUM(amount) as balance FROM transactions WHERE user_id = ?", (user_id,))
        result = cursor.fetchone()
        return result[0] if result and result[0] is not None else 0

    def get_monthly_transactions(self, user_id: int, month: int, year: int):
        cursor = self.db.cursor()
        cursor.execute("""
            SELECT type, SUM(amount) as total, COUNT(*) as count
            FROM transactions
            WHERE user_id = ? AND strftime('%m', created_at) = ? AND strftime('%Y', created_at) = ?
            GROUP BY type
        """, (user_id, f"{month:02d}", str(year)))
        return cursor.fetchall()

    def get_yearly_income(self, user_id: int, year: int) -> float:
        cursor = self.db.cursor()
        cursor.execute("""
            SELECT SUM(amount) as total_income
            FROM transactions
            WHERE user_id = ? AND type = 'income' AND strftime('%Y', created_at) = ?
        """, (user_id, str(year)))
        result = cursor.fetchone()
        return result[0] if result and result[0] is not None else 0

    def get_yearly_deductions(self, user_id: int, year: int) -> float:
        cursor = self.db.cursor()
        cursor.execute("""
            SELECT SUM(amount) as total_deductions
            FROM transactions
            WHERE user_id = ? AND type = 'deduction' AND strftime('%Y', created_at) = ?
        """, (user_id, str(year)))
        result = cursor.fetchone()
        return result[0] if result and result[0] is not None else 0
