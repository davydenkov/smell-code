from abc import ABC, abstractmethod

# 18. Self-Encapsulate Field
#
# BEFORE: Direct field access
class PersonBefore:
    def __init__(self, name):
        self.name = name  # Direct access

    def get_name(self):
        return self.name

    def set_name(self, name):
        self.name = name

# AFTER: Self-encapsulate field
class PersonAfter:
    def __init__(self, name):
        self._name = name

    def get_name(self):
        return self._name

    def set_name(self, name):
        self._name = name

# 19. Replacing the data value with an object (Replace Data Value with Object)
#
# BEFORE: Primitive data type that should be an object
class OrderBefore:
    def __init__(self, customer):
        self._customer = customer  # Just a string

    def get_customer_name(self):
        return self._customer

    def set_customer(self, customer):
        self._customer = customer

# AFTER: Replace with object
class Customer:
    def __init__(self, name):
        self._name = name

    def get_name(self):
        return self._name

class OrderAfter:
    def __init__(self, customer):
        self._customer = customer

    def get_customer(self):
        return self._customer

    def set_customer(self, customer):
        self._customer = customer

    def get_customer_name(self):
        return self._customer.get_name()

# 20. Replacing the value with a reference (Change Value to Reference)
#
# BEFORE: Multiple instances of same object
class CustomerValue:
    def __init__(self, name):
        self._name = name

    def get_name(self):
        return self._name

class OrderValue:
    def __init__(self, customer_name):
        self._customer = CustomerValue(customer_name)  # New instance for each order

# AFTER: Use reference to single instance
class CustomerReference:
    _instances = {}

    def __init__(self, name):
        self._name = name

    @classmethod
    def create(cls, name):
        if name not in cls._instances:
            cls._instances[name] = cls(name)
        return cls._instances[name]

    def get_name(self):
        return self._name

class OrderReference:
    def __init__(self, customer_name):
        self._customer = CustomerReference.create(customer_name)

# 21. Replacing a reference with a value (Change Reference to Value)
#
# BEFORE: Unnecessary reference when value would suffice
class CurrencyReference:
    def __init__(self, code):
        self._code = code

    def get_code(self):
        return self._code

class ProductReference:
    def __init__(self, price, currency):
        self._price = price
        self._currency = currency  # Reference object

# AFTER: Use value object instead
class CurrencyValue:
    def __init__(self, code):
        self._code = code

    def get_code(self):
        return self._code

class ProductValue:
    def __init__(self, price, currency_code):
        self._price = price
        self._currency_code = currency_code  # Just the value

    def get_currency_code(self):
        return self._currency_code

# 22. Replacing an array with an object (Replace Array with Object)
#
# BEFORE: Using array for structured data
class PerformanceArray:
    def get_performance_data(self):
        return {
            'goals': 10,
            'assists': 5,
            'minutes': 120
        }

    def calculate_score(self, data):
        return (data['goals'] * 2) + (data['assists'] * 1.5) + (data['minutes'] / 60)

# AFTER: Replace array with object
class PerformanceData:
    def __init__(self, goals, assists, minutes):
        self._goals = goals
        self._assists = assists
        self._minutes = minutes

    def get_goals(self):
        return self._goals

    def get_assists(self):
        return self._assists

    def get_minutes(self):
        return self._minutes

    def calculate_score(self):
        return (self._goals * 2) + (self._assists * 1.5) + (self._minutes / 60)

class PerformanceObject:
    def get_performance_data(self):
        return PerformanceData(10, 5, 120)

    def calculate_score(self, data):
        return data.calculate_score()

# 23. Duplication of visible data (Duplicate Observed Data)
#
# BEFORE: Domain data mixed with presentation
class OrderDomain:
    def __init__(self):
        self._total = 0

    def add_item(self, price):
        self._total += price
        # Have to update UI here too
        self._update_display()

    def _update_display(self):
        # Update UI elements
        pass

# AFTER: Separate domain and presentation data
class OrderDomainSeparated:
    def __init__(self):
        self._total = 0
        self._observers = []

    def add_item(self, price):
        self._total += price
        self._notify_observers()

    def get_total(self):
        return self._total

    def add_observer(self, observer):
        self._observers.append(observer)

    def _notify_observers(self):
        for observer in self._observers:
            observer.update(self._total)

class OrderDisplay:
    def __init__(self, order):
        self._order = order
        self._order.add_observer(self)

    def update(self, total):
        # Update display with new total
        pass

# 24. Replacing Unidirectional communication with Bidirectional
# communication (Change Unidirectional Association to Bidirectional)
#
# BEFORE: One-way association
class CustomerUni:
    def __init__(self):
        self._orders = []

    def add_order(self, order):
        self._orders.append(order)
        # Order doesn't know about customer

class OrderUni:
    def __init__(self):
        self._items = []

# AFTER: Bidirectional association
class CustomerBi:
    def __init__(self):
        self._orders = []

    def add_order(self, order):
        self._orders.append(order)
        order.set_customer(self)

class OrderBi:
    def __init__(self):
        self._customer = None
        self._items = []

    def set_customer(self, customer):
        self._customer = customer

    def get_customer(self):
        return self._customer

# 25. Replacing Bidirectional communication with Unidirectional
# communication (Change Bidirectional Association to Unidirectional)
#
# BEFORE: Unnecessary bidirectional association
class CustomerBidirectional:
    def __init__(self):
        self._orders = []

    def add_order(self, order):
        self._orders.append(order)
        order.set_customer(self)

class OrderBidirectional:
    def __init__(self):
        self._customer = None

    def set_customer(self, customer):
        self._customer = customer

    def get_customer(self):
        return self._customer

# AFTER: Remove bidirectional link
class CustomerUnidirectional:
    def __init__(self):
        self._orders = []

    def add_order(self, order):
        self._orders.append(order)

class OrderUnidirectional:
    def __init__(self, customer_id):
        self._customer_id = customer_id

    def get_customer_id(self):
        return self._customer_id

# 26. Replacing the magic number with a symbolic constant
# (Replace Magic Number with Symbolic Constant)
#
# BEFORE: Magic numbers
class GeometryBefore:
    def calculate_circle_area(self, radius):
        return 3.14159 * radius * radius  # Magic number

    def calculate_circle_circumference(self, radius):
        return 2 * 3.14159 * radius  # Same magic number

# AFTER: Use symbolic constant
class GeometryAfter:
    PI = 3.14159

    def calculate_circle_area(self, radius):
        return self.PI * radius * radius

    def calculate_circle_circumference(self, radius):
        return 2 * self.PI * radius

# 27. Encapsulate Field
#
# BEFORE: Public field
class PersonPublic:
    def __init__(self, name):
        self.name = name  # Public field

# AFTER: Encapsulated field
class PersonEncapsulated:
    def __init__(self, name):
        self._name = name

    def get_name(self):
        return self._name

    def set_name(self, name):
        self._name = name

# 28. Encapsulate Collection
#
# BEFORE: Direct access to collection
class TeamBefore:
    def __init__(self):
        self.players = []  # Direct access

    def add_player(self, player):
        self.players.append(player)

# AFTER: Encapsulated collection
class TeamAfter:
    def __init__(self):
        self._players = []

    def add_player(self, player):
        self._players.append(player)

    def remove_player(self, player):
        if player in self._players:
            self._players.remove(player)

    def get_players(self):
        return self._players.copy()  # Return copy

    def get_player_count(self):
        return len(self._players)

# 29. Replacing a record with a Data Class
#
# BEFORE: Using array as data structure
class EmployeeArray:
    def create_employee(self, data):
        return {
            'name': data['name'],
            'salary': data['salary'],
            'department': data['department']
        }

    def get_salary(self, employee):
        return employee['salary']

# AFTER: Use data class
class Employee:
    def __init__(self, name, salary, department):
        self._name = name
        self._salary = salary
        self._department = department

    def get_name(self):
        return self._name

    def get_salary(self):
        return self._salary

    def get_department(self):
        return self._department

class EmployeeDataClass:
    def create_employee(self, name, salary, department):
        return Employee(name, salary, department)

    def get_salary(self, employee):
        return employee.get_salary()

# 30. Replacing Type Code with Class
#
# BEFORE: Type code as constants
class EmployeeTypeCode:
    ENGINEER = 0
    SALESMAN = 1
    MANAGER = 2

    def __init__(self, employee_type):
        self._type = employee_type

    def get_type_code(self):
        return self._type

    def get_monthly_salary(self):
        if self._type == self.ENGINEER:
            return 5000
        elif self._type == self.SALESMAN:
            return 4000
        elif self._type == self.MANAGER:
            return 6000
        else:
            return 0

# AFTER: Replace type code with class
class EmployeeType(ABC):
    @abstractmethod
    def get_monthly_salary(self):
        pass

    @staticmethod
    def create_engineer():
        return EngineerType()

    @staticmethod
    def create_salesman():
        return SalesmanType()

    @staticmethod
    def create_manager():
        return ManagerType()

class EngineerType(EmployeeType):
    def get_monthly_salary(self):
        return 5000

class SalesmanType(EmployeeType):
    def get_monthly_salary(self):
        return 4000

class ManagerType(EmployeeType):
    def get_monthly_salary(self):
        return 6000

class EmployeeTypeClass:
    def __init__(self, employee_type):
        self._type = employee_type

    def get_monthly_salary(self):
        return self._type.get_monthly_salary()

# 31. Replacing Type Code with Subclasses
#
# BEFORE: Type code in base class
class EmployeeSubBefore:
    ENGINEER = 0
    SALESMAN = 1
    MANAGER = 2

    def __init__(self, employee_type, salary):
        self._type = employee_type
        self._salary = salary

    def get_salary(self):
        return self._salary

    def get_type(self):
        return self._type

# AFTER: Replace type code with subclasses
class EmployeeSubAfter(ABC):
    def __init__(self, salary):
        self._salary = salary

    def get_salary(self):
        return self._salary

    @abstractmethod
    def get_type(self):
        pass

class Engineer(EmployeeSubAfter):
    def get_type(self):
        return 'engineer'

class Salesman(EmployeeSubAfter):
    def get_type(self):
        return 'salesman'

class Manager(EmployeeSubAfter):
    def get_type(self):
        return 'manager'

# 32. Replacing Type Code with State/Strategy
#
# BEFORE: Type code with behavior
class EmployeeStateBefore:
    JUNIOR = 0
    SENIOR = 1
    LEAD = 2

    def __init__(self, level):
        self._level = level

    def get_salary_multiplier(self):
        if self._level == self.JUNIOR:
            return 1.0
        elif self._level == self.SENIOR:
            return 1.5
        elif self._level == self.LEAD:
            return 2.0
        else:
            return 1.0

# AFTER: Use state/strategy pattern
class EmployeeLevel(ABC):
    @abstractmethod
    def get_salary_multiplier(self):
        pass

class JuniorLevel(EmployeeLevel):
    def get_salary_multiplier(self):
        return 1.0

class SeniorLevel(EmployeeLevel):
    def get_salary_multiplier(self):
        return 1.5

class LeadLevel(EmployeeLevel):
    def get_salary_multiplier(self):
        return 2.0

class EmployeeStateAfter:
    def __init__(self, level):
        self._level = level

    def get_salary_multiplier(self):
        return self._level.get_salary_multiplier()

# 33. Replacing Subclass with Fields
#
# BEFORE: Unnecessary subclasses
class PersonSub(ABC):
    def __init__(self, name, gender):
        self._name = name
        self._gender = gender

    def get_name(self):
        return self._name

    @abstractmethod
    def is_male(self):
        pass

class Male(PersonSub):
    def is_male(self):
        return True

class Female(PersonSub):
    def is_male(self):
        return False

# AFTER: Replace subclass with field
class PersonField:
    def __init__(self, name, gender):
        self._name = name
        self._gender = gender  # 'male' or 'female'

    def get_name(self):
        return self._name

    def is_male(self):
        return self._gender == 'male'

    def get_gender(self):
        return self._gender
