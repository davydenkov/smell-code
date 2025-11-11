import hashlib
import secrets
import sqlite3
from datetime import datetime

class UserManager:
    def __init__(self, db):
        self.db = db

    def register_user(self, user_data):
        # Validate input data
        if not user_data.get('email'):
            raise Exception('Email is required')
        if not user_data.get('password'):
            raise Exception('Password is required')
        if len(user_data['password']) < 8:
            raise Exception('Password must be at least 8 characters')
        if not self._is_valid_email(user_data['email']):
            raise Exception('Invalid email format')

        # Check if user already exists
        cursor = self.db.cursor()
        cursor.execute("SELECT id FROM users WHERE email = ?", (user_data['email'],))
        if cursor.fetchone():
            raise Exception('User already exists')

        # Hash password
        hashed_password = self._hash_password(user_data['password'])

        # Generate verification token
        verification_token = secrets.token_hex(32)

        # Insert user into database
        cursor.execute(
            "INSERT INTO users (email, password, first_name, last_name, verification_token, created_at) VALUES (?, ?, ?, ?, ?, ?)",
            (
                user_data['email'],
                hashed_password,
                user_data['firstName'],
                user_data['lastName'],
                verification_token,
                datetime.now()
            )
        )

        user_id = cursor.lastrowid

        # Create user profile
        cursor.execute(
            "INSERT INTO user_profiles (user_id, phone, address, city, state, zip_code) VALUES (?, ?, ?, ?, ?, ?)",
            (
                user_id,
                user_data.get('phone'),
                user_data.get('address'),
                user_data.get('city'),
                user_data.get('state'),
                user_data.get('zipCode')
            )
        )

        # Send verification email
        subject = 'Please verify your email address'
        message = f"Hello {user_data['firstName']},\n\n"
        message += "Thank you for registering. Please click the link below to verify your email:\n\n"
        message += f"http://example.com/verify?token={verification_token}\n\n"
        message += "Best regards,\nThe Team"

        print(f"Sending email to {user_data['email']}")
        print(f"Subject: {subject}")
        print(f"Message: {message}")

        # Log registration
        log_message = f"User registered: {user_data['email']} at {datetime.now()}"
        print(f"LOG: {log_message}")

        # Create default settings
        cursor.execute(
            "INSERT INTO user_settings (user_id, theme, notifications_enabled) VALUES (?, 'light', 1)",
            (user_id,)
        )

        # Send welcome notification
        cursor.execute(
            "INSERT INTO notifications (user_id, type, message, created_at) VALUES (?, 'welcome', 'Welcome to our platform!', ?)",
            (user_id, datetime.now())
        )

        self.db.commit()
        return user_id

    def _is_valid_email(self, email):
        import re
        pattern = r'^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$'
        return re.match(pattern, email) is not None

    def _hash_password(self, password):
        return hashlib.sha256(password.encode()).hexdigest()
