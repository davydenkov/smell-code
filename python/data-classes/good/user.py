class User:
    def __init__(self, id, name, email, age, email_validator=None):
        self.id = id
        self.name = name
        self.age = age
        self.email_validator = email_validator or EmailValidator()

        self.set_email(email)  # Use setter for validation

    def get_id(self):
        return self.id

    def get_name(self):
        return self.name

    def get_email(self):
        return self.email

    def set_email(self, email):
        if not self.email_validator.is_valid(email):
            raise ValueError('Invalid email address')
        self.email = email

    def get_age(self):
        return self.age

    def set_age(self, age):
        if age < 0 or age > 150:
            raise ValueError('Age must be between 0 and 150')
        self.age = age

    def get_display_name(self):
        return f"{self.name} ({self.age} years old)"

    def can_vote(self):
        return self.age >= 18

    def is_adult(self):
        return self.age >= 18

    def get_age_category(self):
        if self.age < 13:
            return 'child'
        if self.age < 20:
            return 'teenager'
        if self.age < 65:
            return 'adult'
        return 'senior'


class EmailValidator:
    def is_valid(self, email):
        import re
        pattern = r'^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$'
        return re.match(pattern, email) is not None
