class EmailService:
    def send_verification_email(self, email, first_name, verification_token):
        subject = 'Please verify your email address'
        message = f"Hello {first_name},\n\n"
        message += "Thank you for registering. Please click the link below to verify your email:\n\n"
        message += f"http://example.com/verify?token={verification_token}\n\n"
        message += "Best regards,\nThe Team"

        print(f"Sending email to {email}")
        print(f"Subject: {subject}")
        print(f"Message: {message}")


class NotificationService:
    def __init__(self, db):
        self.db = db

    def send_welcome_notification(self, user_id):
        cursor = self.db.cursor()
        cursor.execute(
            "INSERT INTO notifications (user_id, type, message, created_at) VALUES (?, 'welcome', 'Welcome to our platform!', ?)",
            (user_id, datetime.now())
        )
        self.db.commit()


class Logger:
    @staticmethod
    def log_registration(email):
        from datetime import datetime
        log_message = f"User registered: {email} at {datetime.now()}"
        print(f"LOG: {log_message}")
