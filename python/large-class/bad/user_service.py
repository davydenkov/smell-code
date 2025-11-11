import hashlib
import secrets
import sqlite3

class UserService:
    def __init__(self, db, email_config, payment_config):
        self.db = db
        self.email_config = email_config
        self.payment_config = payment_config

        # User properties
        self.user_id = None
        self.user_email = None
        self.user_name = None
        self.user_balance = None

        # Email properties
        self.smtp_host = email_config.get('host')
        self.smtp_port = email_config.get('port')
        self.smtp_username = email_config.get('username')
        self.smtp_password = email_config.get('password')

        # Payment properties
        self.stripe_secret_key = payment_config.get('stripe_secret')
        self.paypal_client_id = payment_config.get('paypal_client_id')
        self.paypal_secret = payment_config.get('paypal_secret')

    # User management methods
    def create_user(self, email, name, password):
        hashed_password = self._hash_password(password)
        cursor = self.db.cursor()
        cursor.execute(
            "INSERT INTO users (email, name, password) VALUES (?, ?, ?)",
            (email, name, hashed_password)
        )
        self.db.commit()
        return cursor.lastrowid

    def authenticate_user(self, email, password):
        cursor = self.db.cursor()
        cursor.execute("SELECT id, password FROM users WHERE email = ?", (email,))
        user = cursor.fetchone()

        if user and self._verify_password(password, user[1]):
            self.user_id = user[0]
            self.user_email = email
            return user[0]
        return False

    def update_user_profile(self, user_id, name, email):
        cursor = self.db.cursor()
        cursor.execute("UPDATE users SET name = ?, email = ? WHERE id = ?", (name, email, user_id))
        self.db.commit()

    def get_user_balance(self, user_id):
        cursor = self.db.cursor()
        cursor.execute("SELECT balance FROM users WHERE id = ?", (user_id,))
        result = cursor.fetchone()
        return result[0] if result else 0

    # Email methods
    def send_welcome_email(self, email, name):
        subject = 'Welcome to our platform!'
        message = f"Hello {name},\n\nWelcome to our platform!"

        # Simulate sending email
        print(f"Sending email to {email} from {self.smtp_username}")
        print(f"Subject: {subject}")
        print(f"Message: {message}")
        return True

    def send_password_reset_email(self, email, reset_token):
        subject = 'Password Reset'
        message = f"Click here to reset your password: http://example.com/reset?token={reset_token}"

        # Simulate sending email
        print(f"Sending email to {email} from {self.smtp_username}")
        print(f"Subject: {subject}")
        print(f"Message: {message}")
        return True

    def send_notification_email(self, email, subject, message):
        # Simulate sending email
        print(f"Sending email to {email} from {self.smtp_username}")
        print(f"Subject: {subject}")
        print(f"Message: {message}")
        return True

    # Payment methods
    def process_stripe_payment(self, amount, token):
        # Simulate Stripe payment processing
        if token and amount > 0:
            return {'success': True, 'transaction_id': f'stripe_{secrets.token_hex(8)}'}
        return {'success': False, 'error': 'Invalid payment data'}

    def process_paypal_payment(self, amount, paypal_token):
        # Simulate PayPal payment processing
        if paypal_token and amount > 0:
            return {'success': True, 'transaction_id': f'paypal_{secrets.token_hex(8)}'}
        return {'success': False, 'error': 'Invalid payment data'}

    def refund_payment(self, transaction_id, amount):
        # Simulate refund logic
        if transaction_id.startswith('stripe_'):
            return {'success': True, 'refund_id': f'refund_{secrets.token_hex(8)}'}
        elif transaction_id.startswith('paypal_'):
            return {'success': True, 'refund_id': f'refund_{secrets.token_hex(8)}'}
        return {'success': False, 'error': 'Unknown transaction type'}

    # Reporting methods
    def generate_user_report(self, user_id):
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

    def generate_sales_report(self, start_date, end_date):
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

    # Utility methods
    def validate_email(self, email):
        import re
        pattern = r'^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$'
        return re.match(pattern, email) is not None

    def generate_token(self):
        return secrets.token_hex(32)

    def _hash_password(self, password):
        return hashlib.sha256(password.encode()).hexdigest()

    def _verify_password(self, password, hashed):
        return self._hash_password(password) == hashed

    def log_activity(self, user_id, action):
        cursor = self.db.cursor()
        cursor.execute(
            "INSERT INTO activity_log (user_id, action, created_at) VALUES (?, ?, datetime('now'))",
            (user_id, action)
        )
        self.db.commit()
