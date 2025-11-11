from datetime import datetime, timedelta
import copy

# 9. Substitution Algorithm
#
# BEFORE: Complex algorithm that can be simplified
class PricingServiceBefore:
    def calculate_price(self, items):
        total = 0
        for item in items:
            if item['type'] == 'book':
                total += item['price'] * 0.9  # 10% discount for books
            elif item['type'] == 'electronics':
                total += item['price'] * 1.1  # 10% markup for electronics
            else:
                total += item['price']
        return total

# AFTER: Substitute with a simpler algorithm
class PricingServiceAfter:
    def __init__(self):
        self._discounts = {
            'book': 0.9,
            'electronics': 1.1,
            'default': 1.0
        }

    def calculate_price(self, items):
        total = 0
        for item in items:
            multiplier = self._discounts.get(item['type'], self._discounts['default'])
            total += item['price'] * multiplier
        return total

# 10. Moving functions between objects (Move Method)
#
# BEFORE: Method in wrong class
class AccountBefore:
    def __init__(self, balance):
        self._balance = balance

    def get_balance(self):
        return self._balance

    # This method belongs in Bank class, not Account
    def transfer_to(self, target, amount):
        if self._balance >= amount:
            self._balance -= amount
            target._balance += amount
            return True
        return False

# AFTER: Move method to appropriate class
class AccountAfter:
    def __init__(self, balance):
        self._balance = balance

    def get_balance(self):
        return self._balance

    def decrease_balance(self, amount):
        self._balance -= amount

    def increase_balance(self, amount):
        self._balance += amount

class Bank:
    def transfer(self, from_account, to_account, amount):
        if from_account.get_balance() >= amount:
            from_account.decrease_balance(amount)
            to_account.increase_balance(amount)
            return True
        return False

# 11. Moving the field (Move Field)
#
# BEFORE: Field in wrong class
class CustomerBefore:
    def __init__(self, name, street, city, zip_code):
        self._name = name
        self._address = {'street': street, 'city': city, 'zip_code': zip_code}

    def get_address(self):
        return f"{self._address['street']}, {self._address['city']} {self._address['zip_code']}"

# AFTER: Move field to dedicated class
class Address:
    def __init__(self, street, city, zip_code):
        self._street = street
        self._city = city
        self._zip_code = zip_code

    def get_full_address(self):
        return f"{self._street}, {self._city} {self._zip_code}"

class CustomerAfter:
    def __init__(self, name, address):
        self._name = name
        self._address = address

    def get_address(self):
        return self._address.get_full_address()

# 12. Class Allocation (Extract Class)
#
# BEFORE: Class has too many responsibilities
class PersonBefore:
    def __init__(self, name, phone_number, office_area_code, office_number):
        self._name = name
        self._phone_number = phone_number
        self._office_area_code = office_area_code
        self._office_number = office_number

    def get_telephone_number(self):
        return f"({self._office_area_code}) {self._office_number}"

# AFTER: Extract telephone number to separate class
class TelephoneNumber:
    def __init__(self, area_code, number):
        self._area_code = area_code
        self._number = number

    def get_telephone_number(self):
        return f"({self._area_code}) {self._number}"

class PersonAfter:
    def __init__(self, name):
        self._name = name
        self._office_telephone = None

    def get_office_telephone(self):
        return self._office_telephone.get_telephone_number() if self._office_telephone else None

    def set_office_telephone(self, telephone):
        self._office_telephone = telephone

# 13. Embedding a class (Inline Class)
#
# BEFORE: Unnecessary class with single responsibility
class OrderValidator:
    def is_valid(self, order):
        return order['total'] > 0

class OrderProcessorBefore:
    def __init__(self):
        self._validator = OrderValidator()

    def process(self, order):
        if self._validator.is_valid(order):
            # Process order
            pass

# AFTER: Inline the class
class OrderProcessorAfter:
    def process(self, order):
        if self._is_valid_order(order):
            # Process order
            pass

    def _is_valid_order(self, order):
        return order['total'] > 0

# 14. Hiding delegation (Hide Delegate)
#
# BEFORE: Client has to know about delegation
class DepartmentBefore:
    def __init__(self, manager):
        self._manager = manager

    def get_manager(self):
        return self._manager

class PersonBefore:
    def __init__(self, department):
        self._department = department

    def get_department(self):
        return self._department

# Client code would be: manager = person.get_department().get_manager()

# AFTER: Hide the delegation
class DepartmentAfter:
    def __init__(self, manager):
        self._manager = manager

    def get_manager(self):
        return self._manager

class PersonAfter:
    def __init__(self, department):
        self._department = department

    def get_department(self):
        return self._department

    def get_manager(self):
        return self._department.get_manager()

# Client code is now: manager = person.get_manager()

# 15. Removing the intermediary (Remove Middle Man)
#
# BEFORE: Too much delegation
class PersonWithMiddleMan:
    def __init__(self, department):
        self._department = department

    def get_department(self):
        return self._department

    def get_manager(self):
        return self._department.get_manager()

    def get_department_name(self):
        return self._department.get_name()

# AFTER: Remove middle man if delegation is too heavy
class PersonDirect:
    def __init__(self, department, manager):
        self._department = department
        self._manager = manager  # Direct reference

    def get_manager(self):
        return self._manager

    def get_department(self):
        return self._department

# 16. Introduction of an external method (Introduce Foreign Method)
#
# BEFORE: Using external class method in wrong place
class ReportGeneratorBefore:
    def generate_report(self):
        date = datetime.now()
        next_month = date.replace(month=date.month + 1)  # Foreign method usage

        # Generate report for next month
        pass

# AFTER: Introduce foreign method
class ReportGeneratorAfter:
    def generate_report(self):
        date = datetime.now()
        next_month = self._next_month(date)

        # Generate report for next month
        pass

    def _next_month(self, date):
        # Handle year rollover
        if date.month == 12:
            return date.replace(year=date.year + 1, month=1)
        else:
            return date.replace(month=date.month + 1)

# 17. The introduction of local extension (Introduce Local Extension)
#
# BEFORE: Adding methods to external class (not possible in Python)

# AFTER: Create local extension class
class DateTimeExtension(datetime):
    def next_month(self):
        if self.month == 12:
            return self.replace(year=self.year + 1, month=1)
        else:
            return self.replace(month=self.month + 1)

    def previous_month(self):
        if self.month == 1:
            return self.replace(year=self.year - 1, month=12)
        else:
            return self.replace(month=self.month - 1)
