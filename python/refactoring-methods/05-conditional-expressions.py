import re
from abc import ABC, abstractmethod
from datetime import datetime

# 34. Decomposition of a conditional operator (Decompose Conditional)
#
# BEFORE: Complex conditional logic
class PaymentProcessorBefore:
    def calculate_fee(self, amount, is_international, is_premium):
        fee = 0
        if amount > 100 and is_international and is_premium:
            fee = amount * 0.05 + 10
        elif amount > 100 and is_international and not is_premium:
            fee = amount * 0.05 + 15
        elif amount <= 100 and is_international:
            fee = amount * 0.03 + 5
        else:
            fee = amount * 0.02
        return fee

# AFTER: Decompose conditional
class PaymentProcessorAfter:
    def calculate_fee(self, amount, is_international, is_premium):
        if self._is_high_value_international_premium(amount, is_international, is_premium):
            return self._calculate_high_value_international_premium_fee(amount)
        elif self._is_high_value_international_standard(amount, is_international, is_premium):
            return self._calculate_high_value_international_standard_fee(amount)
        elif self._is_low_value_international(amount, is_international):
            return self._calculate_low_value_international_fee(amount)
        else:
            return self._calculate_domestic_fee(amount)

    def _is_high_value_international_premium(self, amount, is_international, is_premium):
        return amount > 100 and is_international and is_premium

    def _is_high_value_international_standard(self, amount, is_international, is_premium):
        return amount > 100 and is_international and not is_premium

    def _is_low_value_international(self, amount, is_international):
        return amount <= 100 and is_international

    def _calculate_high_value_international_premium_fee(self, amount):
        return amount * 0.05 + 10

    def _calculate_high_value_international_standard_fee(self, amount):
        return amount * 0.05 + 15

    def _calculate_low_value_international_fee(self, amount):
        return amount * 0.03 + 5

    def _calculate_domestic_fee(self, amount):
        return amount * 0.02

# 35. Consolidation of a conditional expression (Consolidate Conditional Expression)
#
# BEFORE: Multiple conditionals with same result
class InsuranceCalculatorBefore:
    def is_eligible_for_discount(self, age, is_student, has_good_record):
        if age < 25:
            return False
        if is_student:
            return True
        if has_good_record:
            return True
        return False

# AFTER: Consolidate conditionals
class InsuranceCalculatorAfter:
    def is_eligible_for_discount(self, age, is_student, has_good_record):
        return age >= 25 and (is_student or has_good_record)

# 36. Consolidation of duplicate conditional fragments
# (Consolidate Duplicate Conditional Fragments)
#
# BEFORE: Duplicate code in conditional branches
class FileProcessorBefore:
    def process_file(self, file_path):
        if self._is_valid_file(file_path):
            self._log_processing(file_path)
            self._validate_content(file_path)
            self._save_to_database(file_path)
            self._send_notification(file_path)
        else:
            self._log_error(file_path)
            self._send_notification(file_path)  # Duplicate

    def _send_notification(self, file_path):
        # Send notification logic
        pass

# AFTER: Consolidate duplicate fragments
class FileProcessorAfter:
    def process_file(self, file_path):
        self._send_notification(file_path)  # Moved outside conditional

        if self._is_valid_file(file_path):
            self._log_processing(file_path)
            self._validate_content(file_path)
            self._save_to_database(file_path)
        else:
            self._log_error(file_path)

    def _send_notification(self, file_path):
        # Send notification logic
        pass

# 37. Remove Control Flag
#
# BEFORE: Control flag to break out of loop
class DataProcessorBefore:
    def find_person(self, people, name):
        found = False
        for person in people:
            if person['name'] == name:
                found = person
                break  # Control flag usage
        return found

# AFTER: Remove control flag
class DataProcessorAfter:
    def find_person(self, people, name):
        for person in people:
            if person['name'] == name:
                return person  # Direct return
        return False

# 38. Replacing Nested Conditional statements with a boundary operator
# (Replace Nested Conditional with Guard Clauses)
#
# BEFORE: Nested conditionals
class PaymentValidatorBefore:
    def is_valid_payment(self, payment):
        if payment['amount'] > 0:
            if payment['card_number'] is not None:
                if len(payment['card_number']) == 16:
                    if self._is_valid_expiry(payment['expiry']):
                        return True
        return False

    def _is_valid_expiry(self, expiry):
        return datetime.strptime(expiry, '%Y-%m-%d') > datetime.now()

# AFTER: Replace with guard clauses
class PaymentValidatorAfter:
    def is_valid_payment(self, payment):
        if payment['amount'] <= 0:
            return False

        if payment['card_number'] is None:
            return False

        if len(payment['card_number']) != 16:
            return False

        if not self._is_valid_expiry(payment['expiry']):
            return False

        return True

    def _is_valid_expiry(self, expiry):
        return datetime.strptime(expiry, '%Y-%m-%d') > datetime.now()

# 39. Replacing a conditional operator with polymorphism (Replace Conditional with Polymorphism)
#
# BEFORE: Type checking with conditionals
class BirdBefore:
    EUROPEAN = 'european'
    AFRICAN = 'african'
    NORWEGIAN_BLUE = 'norwegian_blue'

    def __init__(self, bird_type):
        self._type = bird_type

    def get_speed(self):
        if self._type == self.EUROPEAN:
            return self._get_base_speed()
        elif self._type == self.AFRICAN:
            return self._get_base_speed() - self._voltage * 2
        elif self._type == self.NORWEGIAN_BLUE:
            return 0 if self._is_nailed else self._get_base_speed()
        else:
            return self._get_base_speed()

    def _get_base_speed(self):
        return 10

# AFTER: Replace conditional with polymorphism
class BirdAfter(ABC):
    @abstractmethod
    def get_speed(self):
        pass

    def _get_base_speed(self):
        return 10

class EuropeanSwallow(BirdAfter):
    def get_speed(self):
        return self._get_base_speed()

class AfricanSwallow(BirdAfter):
    def __init__(self, voltage):
        self._voltage = voltage

    def get_speed(self):
        return self._get_base_speed() - self._voltage * 2

class NorwegianBlueParrot(BirdAfter):
    def __init__(self, is_nailed):
        self._is_nailed = is_nailed

    def get_speed(self):
        return 0 if self._is_nailed else self._get_base_speed()

# 40. Introduction of the object (Introduce Object)
#
# BEFORE: Primitive obsession with conditionals
class UserValidatorBefore:
    def validate_user(self, user):
        if not user['name']:
            return 'Name is required'

        if len(user['name']) < 2:
            return 'Name must be at least 2 characters'

        if not self._is_valid_email(user['email']):
            return 'Invalid email format'

        return True

    def _is_valid_email(self, email):
        # Simple email validation
        return '@' in email and '.' in email

# AFTER: Introduce validation result object
class ValidationResult:
    def __init__(self, is_valid=True, errors=None):
        self._is_valid = is_valid
        self._errors = errors or []

    def is_valid(self):
        return self._is_valid

    def get_errors(self):
        return self._errors

    def add_error(self, error):
        self._is_valid = False
        self._errors.append(error)
        return self

class UserValidatorAfter:
    def validate_user(self, user):
        result = ValidationResult()

        if not user['name']:
            result.add_error('Name is required')

        if len(user['name']) < 2:
            result.add_error('Name must be at least 2 characters')

        if not self._is_valid_email(user['email']):
            result.add_error('Invalid email format')

        return result

    def _is_valid_email(self, email):
        # Simple email validation
        return '@' in email and '.' in email

# 41. Introduction of the statement (Introduction Statement)
#
# BEFORE: Magic assertion
class AccountAssertion:
    def __init__(self, balance):
        self._balance = balance

    def withdraw(self, amount):
        assert amount > 0 and amount <= self._balance
        self._balance -= amount

# AFTER: Introduce assertion method
class AccountAssertionAfter:
    def __init__(self, balance):
        self._balance = balance

    def withdraw(self, amount):
        self._assert_valid_withdrawal(amount)
        self._balance -= amount

    def _assert_valid_withdrawal(self, amount):
        assert amount > 0 and amount <= self._balance, 'Invalid withdrawal amount'
