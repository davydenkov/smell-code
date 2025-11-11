import re

class UserValidator:
    def validate_registration_data(self, user_data):
        if not user_data.get('email'):
            raise Exception('Email is required')
        if not user_data.get('password'):
            raise Exception('Password is required')
        if len(user_data['password']) < 8:
            raise Exception('Password must be at least 8 characters')
        if not self._is_valid_email(user_data['email']):
            raise Exception('Invalid email format')

    def _is_valid_email(self, email):
        pattern = r'^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$'
        return re.match(pattern, email) is not None
