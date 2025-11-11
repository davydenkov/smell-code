class UserManager:
    def __init__(self, db):
        self.validator = UserValidator()
        self.repository = UserRepository(db)
        self.email_service = EmailService()
        self.notification_service = NotificationService(db)

    def register_user(self, user_data):
        self.validator.validate_registration_data(user_data)

        if self.repository.user_exists(user_data['email']):
            raise Exception('User already exists')

        user_data = self._prepare_user_data(user_data)

        user_id = self.repository.create_user(user_data)
        self.repository.create_user_profile(user_id, user_data)
        self.repository.create_user_settings(user_id)

        self.email_service.send_verification_email(
            user_data['email'],
            user_data['firstName'],
            user_data['verificationToken']
        )

        self.notification_service.send_welcome_notification(user_id)
        Logger.log_registration(user_data['email'])

        return user_id

    def _prepare_user_data(self, user_data):
        import hashlib
        import secrets

        user_data['password'] = hashlib.sha256(user_data['password'].encode()).hexdigest()
        user_data['verificationToken'] = secrets.token_hex(32)
        return user_data
