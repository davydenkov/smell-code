class EmailService:
    def __init__(self, email_config):
        self.smtp_username = email_config.get('username')

    def send_welcome_email(self, email: str, name: str) -> bool:
        subject = 'Welcome to our platform!'
        message = f"Hello {name},\n\nWelcome to our platform!"

        print(f"Sending email to {email} from {self.smtp_username}")
        print(f"Subject: {subject}")
        print(f"Message: {message}")
        return True

    def send_password_reset_email(self, email: str, reset_token: str) -> bool:
        subject = 'Password Reset'
        message = f"Click here to reset your password: http://example.com/reset?token={reset_token}"

        print(f"Sending email to {email} from {self.smtp_username}")
        print(f"Subject: {subject}")
        print(f"Message: {message}")
        return True

    def send_notification_email(self, email: str, subject: str, message: str) -> bool:
        print(f"Sending email to {email} from {self.smtp_username}")
        print(f"Subject: {subject}")
        print(f"Message: {message}")
        return True
