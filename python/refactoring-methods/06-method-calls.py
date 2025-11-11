# 42. Renaming a method (Rename Method)
#
# BEFORE: Poorly named method
class CalculatorBefore:
    def calc(self, a, b):  # Unclear name
        return a + b

# AFTER: Rename method to be more descriptive
class CalculatorAfter:
    def add(self, a, b):
        return a + b

# 43. Adding a parameter (Add Parameter)
#
# BEFORE: Method missing required parameter
class EmailSenderBefore:
    def send_email(self, to_email, subject, body):
        # Send email with default priority
        priority = 'normal'
        # Send logic
        pass

# AFTER: Add parameter
class EmailSenderAfter:
    def send_email(self, to_email, subject, body, priority='normal'):
        # Send logic with priority
        pass

# 44. Deleting a parameter (Remove Parameter)
#
# BEFORE: Unnecessary parameter
class ReportGeneratorBefore:
    def generate_report(self, data, report_format, include_header=True):
        if report_format == 'html':
            # Always include header for HTML
            include_header = True
        # Generate report
        pass

# AFTER: Remove unnecessary parameter
class ReportGeneratorAfter:
    def generate_report(self, data, report_format):
        include_header = (report_format == 'html')
        # Generate report
        pass

# 45. Separation of Query and Modifier (Separate Query from Modifier)
#
# BEFORE: Method that both queries and modifies
class BankAccountBefore:
    def __init__(self):
        self._balance = 0

    def withdraw(self, amount):
        if self._balance >= amount:
            self._balance -= amount
            return True
        return False

# AFTER: Separate query from modifier
class BankAccountAfter:
    def __init__(self):
        self._balance = 0

    def can_withdraw(self, amount):
        return self._balance >= amount

    def withdraw(self, amount):
        if self.can_withdraw(amount):
            self._balance -= amount
            return True
        return False

# 46. Parameterization of the method (Parameterize Method)
#
# BEFORE: Similar methods with different values
class ReportGeneratorParamBefore:
    def generate_weekly_report(self):
        return self.generate_report(7)

    def generate_monthly_report(self):
        return self.generate_report(30)

    def generate_quarterly_report(self):
        return self.generate_report(90)

    def generate_report(self, days):
        # Generate report for specified days
        pass

# AFTER: Parameterize method
class ReportGeneratorParamAfter:
    def generate_report(self, days):
        # Generate report for specified days
        pass

    def generate_weekly_report(self):
        return self.generate_report(7)

    def generate_monthly_report(self):
        return self.generate_report(30)

    def generate_quarterly_report(self):
        return self.generate_report(90)

# 47. Replacing a parameter with explicit methods (Replace Parameter with Explicit Methods)
#
# BEFORE: Parameter determines behavior
class EmployeeExplicitBefore:
    ENGINEER = 0
    SALESMAN = 1
    MANAGER = 2

    def __init__(self, employee_type):
        self._type = employee_type

    def get_salary(self, base_salary):
        if self._type == self.ENGINEER:
            return base_salary * 1.0
        elif self._type == self.SALESMAN:
            return base_salary * 1.1
        elif self._type == self.MANAGER:
            return base_salary * 1.2
        else:
            return base_salary

# AFTER: Replace parameter with explicit methods
class EmployeeExplicitAfter:
    def get_engineer_salary(self, base_salary):
        return base_salary * 1.0

    def get_salesman_salary(self, base_salary):
        return base_salary * 1.1

    def get_manager_salary(self, base_salary):
        return base_salary * 1.2

# 48. Save the Whole Object
#
# BEFORE: Passing individual fields
class OrderWholeBefore:
    def __init__(self, customer_name, customer_address):
        self._customer = {'name': customer_name, 'address': customer_address}

    def calculate_shipping(self):
        return self._get_shipping_cost(self._customer['name'], self._customer['address'])

    def _get_shipping_cost(self, name, address):
        # Calculate based on name and address
        return 10.0

# AFTER: Pass whole object
class Customer:
    def __init__(self, name, address):
        self._name = name
        self._address = address

    def get_name(self):
        return self._name

    def get_address(self):
        return self._address

class OrderWholeAfter:
    def __init__(self, customer):
        self._customer = customer

    def calculate_shipping(self):
        return self._get_shipping_cost(self._customer)

    def _get_shipping_cost(self, customer):
        # Calculate based on customer object
        return 10.0

# 49. Replacing a parameter with a method call (Replace Parameter with Method)
#
# BEFORE: Parameter calculated outside method
class DiscountCalculatorParamBefore:
    def calculate_discount(self, price, customer_type):
        # customer_type passed in
        return price * self._get_discount_rate(customer_type)

    def _get_discount_rate(self, customer_type):
        if customer_type == 'premium':
            return 0.1
        elif customer_type == 'regular':
            return 0.05
        else:
            return 0.0

class OrderParamBefore:
    def __init__(self, customer):
        self._customer = customer

    def get_discounted_price(self, price):
        calculator = DiscountCalculatorParamBefore()
        return calculator.calculate_discount(price, self._customer.get_type())

# AFTER: Replace parameter with method call
class DiscountCalculatorParamAfter:
    def calculate_discount(self, price, customer):
        return price * self._get_discount_rate(customer.get_type())

    def _get_discount_rate(self, customer_type):
        if customer_type == 'premium':
            return 0.1
        elif customer_type == 'regular':
            return 0.05
        else:
            return 0.0

class OrderParamAfter:
    def __init__(self, customer):
        self._customer = customer

    def get_discounted_price(self, price):
        calculator = DiscountCalculatorParamAfter()
        return calculator.calculate_discount(price, self._customer)

# 50. Introduction of the boundary object (Introduce Parameter Object)
#
# BEFORE: Multiple parameters
class TemperatureRangeBefore:
    def within_range(self, min_temp, max_temp, current_temp):
        return current_temp >= min_temp and current_temp <= max_temp

    def get_average_temp(self, min_temp, max_temp):
        return (min_temp + max_temp) / 2

# AFTER: Introduce parameter object
class TemperatureRange:
    def __init__(self, min_temp, max_temp):
        self._min_temp = min_temp
        self._max_temp = max_temp

    def get_min_temp(self):
        return self._min_temp

    def get_max_temp(self):
        return self._max_temp

    def within_range(self, current_temp):
        return current_temp >= self._min_temp and current_temp <= self._max_temp

    def get_average_temp(self):
        return (self._min_temp + self._max_temp) / 2

# 51. Removing the Value Setting Method
#
# BEFORE: Setter that's not needed
class SensorBefore:
    def __init__(self, temperature):
        self._temperature = temperature

    def get_temperature(self):
        return self._temperature

    def set_temperature(self, temperature):  # Not needed if immutable
        self._temperature = temperature

# AFTER: Remove setter for immutable object
class SensorAfter:
    def __init__(self, temperature):
        self._temperature = temperature

    def get_temperature(self):
        return self._temperature

    # set_temperature removed

# 52. Hiding a method (Hide Method)
#
# BEFORE: Public method that should be private
class DataProcessorHideBefore:
    def validate_data(self, data):  # Should be private
        return data and isinstance(data, list)

    def process_data(self, data):
        if self.validate_data(data):
            # Process data
            pass

# AFTER: Hide method
class DataProcessorHideAfter:
    def _validate_data(self, data):
        return data and isinstance(data, list)

    def process_data(self, data):
        if self._validate_data(data):
            # Process data
            pass

# 53. Replacing the constructor with the factory method (Replace Constructor with Factory Method)
#
# BEFORE: Complex constructor
class ComplexObjectBefore:
    def __init__(self, obj_type, config=None):
        if config is None:
            config = {}
        self._type = obj_type
        self._config = config

        if obj_type == 'database':
            self._config.update({'host': 'localhost', 'port': 3306})
        elif obj_type == 'file':
            self._config.update({'path': '/tmp', 'format': 'json'})

# AFTER: Replace constructor with factory method
class ComplexObjectAfter:
    def __init__(self, obj_type, config):
        self._type = obj_type
        self._config = config

    @staticmethod
    def create_database_connection(config=None):
        if config is None:
            config = {}
        final_config = {'host': 'localhost', 'port': 3306}
        final_config.update(config)
        return ComplexObjectAfter('database', final_config)

    @staticmethod
    def create_file_handler(config=None):
        if config is None:
            config = {}
        final_config = {'path': '/tmp', 'format': 'json'}
        final_config.update(config)
        return ComplexObjectAfter('file', final_config)

# 54. Encapsulation of top-down type conversion (Encapsulate Downcast)
#
# BEFORE: Downcast in client code
class ShapeCollectionBefore:
    def __init__(self):
        self._shapes = []

    def add_shape(self, shape):
        self._shapes.append(shape)

    def get_shapes(self):
        return self._shapes

# Client code would do:
# circles = [shape for shape in collection.get_shapes() if isinstance(shape, Circle)]

# AFTER: Encapsulate downcast
class ShapeCollectionAfter:
    def __init__(self):
        self._shapes = []

    def add_shape(self, shape):
        self._shapes.append(shape)

    def get_circles(self):
        return [shape for shape in self._shapes if isinstance(shape, Circle)]

    def get_squares(self):
        return [shape for shape in self._shapes if isinstance(shape, Square)]

# Placeholder classes for the example
class Circle:
    pass

class Square:
    pass

# 55. Replacing the error code with an exceptional situation (Replace Error Code with Exception)
#
# BEFORE: Error codes
class FileReaderErrorBefore:
    FILE_NOT_FOUND = 1
    PERMISSION_DENIED = 2

    def read_file(self, filename):
        if not os.path.exists(filename):
            return self.FILE_NOT_FOUND

        if not os.access(filename, os.R_OK):
            return self.PERMISSION_DENIED

        with open(filename, 'r') as f:
            return f.read()

# Client code would check for error codes

# AFTER: Replace error codes with exceptions
class FileNotFoundException(Exception):
    pass

class PermissionDeniedException(Exception):
    pass

class FileReaderExceptionAfter:
    def read_file(self, filename):
        if not os.path.exists(filename):
            raise FileNotFoundException(f"File not found: {filename}")

        if not os.access(filename, os.R_OK):
            raise PermissionDeniedException(f"Permission denied: {filename}")

        with open(filename, 'r') as f:
            return f.read()

# Client code would use try/except

# 56. Replacing an exceptional situation with a check (Replace Exception with Test)
#
# BEFORE: Using exception for control flow
class EmptyStackException(Exception):
    pass

class StackExceptionBefore:
    def __init__(self):
        self._items = []

    def pop(self):
        if not self._items:
            raise EmptyStackException()
        return self._items.pop()

# Client code would catch exception

# AFTER: Replace exception with test
class StackTestAfter:
    def __init__(self):
        self._items = []

    def is_empty(self):
        return not self._items

    def pop(self):
        return self._items.pop()

# Client code would check is_empty() first

import os
