class UserRepository:
    def __init__(self, db):
        self.db = db

    def update_user_profile(self, user_id: int, name: str, email: str) -> None:
        cursor = self.db.cursor()
        cursor.execute("UPDATE users SET name = ?, email = ? WHERE id = ?", (name, email, user_id))
        self.db.commit()

    def get_user_email(self, user_id: int) -> str:
        cursor = self.db.cursor()
        cursor.execute("SELECT email FROM users WHERE id = ?", (user_id,))
        result = cursor.fetchone()
        return result[0] if result else None
