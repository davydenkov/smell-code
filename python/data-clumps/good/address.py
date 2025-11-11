class Address:
    def __init__(self, street: str, city: str, state: str, zip_code: str):
        self.street = street
        self.city = city
        self.state = state
        self.zip_code = zip_code

    def get_street(self) -> str:
        return self.street

    def get_city(self) -> str:
        return self.city

    def get_state(self) -> str:
        return self.state

    def get_zip_code(self) -> str:
        return self.zip_code

    def is_valid(self) -> bool:
        if not self.street or not self.city or not self.state or not self.zip_code:
            return False

        if len(self.zip_code) != 5:
            return False

        return True

    def to_string(self) -> str:
        return f"{self.street}, {self.city}, {self.state} {self.zip_code}"

    def to_label_format(self) -> str:
        return f"{self.street}\n{self.city}, {self.state} {self.zip_code}"


class Person:
    def __init__(self, first_name: str, last_name: str, email: str, phone: str = None, date_of_birth: str = None):
        self.first_name = first_name
        self.last_name = last_name
        self.email = email
        self.phone = phone
        self.date_of_birth = date_of_birth

    def get_first_name(self) -> str:
        return self.first_name

    def get_last_name(self) -> str:
        return self.last_name

    def get_full_name(self) -> str:
        return f"{self.first_name} {self.last_name}"

    def get_email(self) -> str:
        return self.email

    def get_phone(self) -> str:
        return self.phone

    def get_date_of_birth(self) -> str:
        return self.date_of_birth

    def is_valid(self) -> bool:
        if not self.first_name or not self.last_name:
            return False
        if not self._is_valid_email(self.email):
            return False
        return True

    def _is_valid_email(self, email: str) -> bool:
        import re
        pattern = r'^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$'
        return re.match(pattern, email) is not None
