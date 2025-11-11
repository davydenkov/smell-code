import hashlib
import secrets

class Utility:
    @staticmethod
    def hash_password(password: str) -> str:
        return hashlib.sha256(password.encode()).hexdigest()

    @staticmethod
    def verify_password(password: str, hashed: str) -> bool:
        return Utility.hash_password(password) == hashed

    @staticmethod
    def validate_email(email: str) -> bool:
        import re
        pattern = r'^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$'
        return re.match(pattern, email) is not None

    @staticmethod
    def generate_token() -> str:
        return secrets.token_hex(32)


class ActivityLogger:
    def __init__(self, db):
        self.db = db

    def log_activity(self, user_id: int, action: str) -> None:
        cursor = self.db.cursor()
        cursor.execute(
            "INSERT INTO activity_log (user_id, action, created_at) VALUES (?, ?, datetime('now'))",
            (user_id, action)
        )
        self.db.commit()
