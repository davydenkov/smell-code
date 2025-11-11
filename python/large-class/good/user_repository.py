import hashlib

class UserRepository:
    def __init__(self, db):
        self.db = db

    def create(self, email: str, name: str, hashed_password: str) -> int:
        cursor = self.db.cursor()
        cursor.execute(
            "INSERT INTO users (email, name, password) VALUES (?, ?, ?)",
            (email, name, hashed_password)
        )
        self.db.commit()
        return cursor.lastrowid

    def authenticate(self, email: str, password: str):
        cursor = self.db.cursor()
        cursor.execute("SELECT id, email, name, balance, password FROM users WHERE email = ?", (email,))
        result = cursor.fetchone()

        if result and Utility.verify_password(password, result[4]):
            return User(result[0], result[1], result[2], result[3])
        return None

    def find_by_id(self, user_id: int):
        cursor = self.db.cursor()
        cursor.execute("SELECT id, email, name, balance FROM users WHERE id = ?", (user_id,))
        result = cursor.fetchone()

        if result:
            return User(result[0], result[1], result[2], result[3])
        return None

    def update(self, user) -> None:
        cursor = self.db.cursor()
        cursor.execute(
            "UPDATE users SET name = ?, email = ? WHERE id = ?",
            (user.get_name(), user.get_email(), user.get_id())
        )
        self.db.commit()
