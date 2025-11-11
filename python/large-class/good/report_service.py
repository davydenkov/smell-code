class ReportService:
    def __init__(self, db):
        self.db = db

    def generate_user_report(self, user_id: int):
        cursor = self.db.cursor()
        cursor.execute("""
            SELECT u.name, u.email, u.created_at,
                   COUNT(o.id) as order_count,
                   SUM(o.total) as total_spent
            FROM users u
            LEFT JOIN orders o ON u.id = o.user_id
            WHERE u.id = ?
            GROUP BY u.id
        """, (user_id,))
        return cursor.fetchone()

    def generate_sales_report(self, start_date: str, end_date: str):
        cursor = self.db.cursor()
        cursor.execute("""
            SELECT DATE(created_at) as date,
                   COUNT(*) as order_count,
                   SUM(total) as total_sales
            FROM orders
            WHERE created_at BETWEEN ? AND ?
            GROUP BY DATE(created_at)
        """, (start_date, end_date))
        return cursor.fetchall()
