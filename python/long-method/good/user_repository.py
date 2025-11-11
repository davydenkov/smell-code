from datetime import datetime

class UserRepository:
    def __init__(self, db):
        self.db = db

    def user_exists(self, email):
        cursor = self.db.cursor()
        cursor.execute("SELECT id FROM users WHERE email = ?", (email,))
        return cursor.fetchone() is not None

    def create_user(self, user_data):
        cursor = self.db.cursor()
        cursor.execute(
            "INSERT INTO users (email, password, first_name, last_name, verification_token, created_at) VALUES (?, ?, ?, ?, ?, ?)",
            (
                user_data['email'],
                user_data['password'],
                user_data['firstName'],
                user_data['lastName'],
                user_data['verificationToken'],
                datetime.now()
            )
        )
        self.db.commit()
        return cursor.lastrowid

    def create_user_profile(self, user_id, user_data):
        cursor = self.db.cursor()
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
        self.db.commit()

    def create_user_settings(self, user_id):
        cursor = self.db.cursor()
        cursor.execute(
            "INSERT INTO user_settings (user_id, theme, notifications_enabled) VALUES (?, 'light', 1)",
            (user_id,)
        )
        self.db.commit()
