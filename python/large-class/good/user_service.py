class UserService:
    def __init__(self, db, email_config, payment_config):
        self.user_repository = UserRepository(db)
        self.email_service = EmailService(email_config)
        self.payment_service = PaymentService(payment_config)
        self.report_service = ReportService(db)
        self.activity_logger = ActivityLogger(db)

    # User management methods
    def create_user(self, email: str, name: str, password: str) -> int:
        hashed_password = Utility.hash_password(password)
        user_id = self.user_repository.create(email, name, hashed_password)

        self.email_service.send_welcome_email(email, name)
        self.activity_logger.log_activity(user_id, 'user_created')

        return user_id

    def authenticate_user(self, email: str, password: str):
        user = self.user_repository.authenticate(email, password)

        if user:
            self.activity_logger.log_activity(user.get_id(), 'user_login')

        return user

    def update_user_profile(self, user_id: int, name: str, email: str) -> None:
        user = self.user_repository.find_by_id(user_id)
        if user:
            user.update_profile(name, email)
            self.user_repository.update(user)
            self.activity_logger.log_activity(user_id, 'profile_updated')

    def get_user_balance(self, user_id: int) -> float:
        user = self.user_repository.find_by_id(user_id)
        return user.get_balance() if user else 0.0

    # Email methods
    def send_welcome_email(self, email: str, name: str) -> bool:
        return self.email_service.send_welcome_email(email, name)

    def send_password_reset_email(self, email: str, reset_token: str) -> bool:
        return self.email_service.send_password_reset_email(email, reset_token)

    def send_notification_email(self, email: str, subject: str, message: str) -> bool:
        return self.email_service.send_notification_email(email, subject, message)

    # Payment methods
    def process_stripe_payment(self, amount: float, token: str):
        return self.payment_service.process_stripe_payment(amount, token)

    def process_paypal_payment(self, amount: float, paypal_token: str):
        return self.payment_service.process_paypal_payment(amount, paypal_token)

    def refund_payment(self, transaction_id: str, amount: float):
        return self.payment_service.refund_payment(transaction_id, amount)

    # Reporting methods
    def generate_user_report(self, user_id: int):
        return self.report_service.generate_user_report(user_id)

    def generate_sales_report(self, start_date: str, end_date: str):
        return self.report_service.generate_sales_report(start_date, end_date)
