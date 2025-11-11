from abc import ABC, abstractmethod

# 69. Separation of inheritance (Tease Apart Inheritance)
#
# BEFORE: Class hierarchy mixing two different responsibilities
class EmployeeTeaseBefore(ABC):
    def __init__(self, name, rate):
        self.name = name
        self.rate = rate

    def get_name(self):
        return self.name

class SalariedEmployeeTeaseBefore(EmployeeTeaseBefore):
    def get_pay(self):
        return self.rate

class CommissionedEmployeeTeaseBefore(EmployeeTeaseBefore):
    def __init__(self, name, rate, commission):
        super().__init__(name, rate)
        self._commission = commission

    def get_pay(self):
        return self.rate + self._commission

# AFTER: Tease apart inheritance into two separate hierarchies
class Payable(ABC):
    @abstractmethod
    def get_pay(self):
        pass

class EmployeeTeaseAfter(ABC):
    def __init__(self, name):
        self.name = name

    def get_name(self):
        return self.name

class SalariedEmployeeTeaseAfter(EmployeeTeaseAfter, Payable):
    def __init__(self, name, salary):
        super().__init__(name)
        self._salary = salary

    def get_pay(self):
        return self._salary

class CommissionedEmployeeTeaseAfter(EmployeeTeaseAfter, Payable):
    def __init__(self, name, base_salary, commission):
        super().__init__(name)
        self._base_salary = base_salary
        self._commission = commission

    def get_pay(self):
        return self._base_salary + self._commission

# 70. Converting a procedural project into objects (Convert Procedural Design to Objects)
#
# BEFORE: Procedural code with global functions and data
class ProceduralDesignBefore:
    _accounts = {}

    @classmethod
    def create_account(cls, account_id, balance):
        cls._accounts[account_id] = balance

    @classmethod
    def get_balance(cls, account_id):
        return cls._accounts.get(account_id, 0)

    @classmethod
    def deposit(cls, account_id, amount):
        if account_id in cls._accounts:
            cls._accounts[account_id] += amount

    @classmethod
    def withdraw(cls, account_id, amount):
        if account_id in cls._accounts and cls._accounts[account_id] >= amount:
            cls._accounts[account_id] -= amount
            return True
        return False

# AFTER: Convert to object-oriented design
class Account:
    def __init__(self, account_id, balance=0):
        self._id = account_id
        self._balance = balance

    def get_id(self):
        return self._id

    def get_balance(self):
        return self._balance

    def deposit(self, amount):
        self._balance += amount

    def withdraw(self, amount):
        if self._balance >= amount:
            self._balance -= amount
            return True
        return False

class Bank:
    def __init__(self):
        self._accounts = {}

    def create_account(self, account_id, balance=0):
        account = Account(account_id, balance)
        self._accounts[account_id] = account
        return account

    def get_account(self, account_id):
        return self._accounts.get(account_id)

    def get_balance(self, account_id):
        account = self.get_account(account_id)
        return account.get_balance() if account else 0

    def deposit(self, account_id, amount):
        account = self.get_account(account_id)
        if account:
            account.deposit(amount)

    def withdraw(self, account_id, amount):
        account = self.get_account(account_id)
        return account.withdraw(amount) if account else False

# 71. Separating the domain from the representation (Separate Domain from Presentation)
#
# BEFORE: Domain logic mixed with presentation
class OrderPresentationBefore:
    def __init__(self):
        self._items = []
        self._total = 0

    def add_item(self, name, price, quantity):
        self._items.append({'name': name, 'price': price, 'quantity': quantity})
        self._total += price * quantity

        # Presentation logic mixed in
        print(f"Added {quantity} x {name} to order")
        print(f"Current total: ${self._total:.2f}")

    def get_total(self):
        return self._total

    def display_order(self):
        print("Order Summary:")
        for item in self._items:
            print(f"- {item['quantity']} x {item['name']} @ ${item['price']:.2f}")
        print(f"Total: ${self._total:.2f}")

# AFTER: Separate domain from presentation
class OrderItem:
    def __init__(self, name, price, quantity):
        self._name = name
        self._price = price
        self._quantity = quantity

    def get_name(self):
        return self._name

    def get_price(self):
        return self._price

    def get_quantity(self):
        return self._quantity

    def get_total(self):
        return self._price * self._quantity

class OrderDomainAfter:
    def __init__(self):
        self._items = []

    def add_item(self, name, price, quantity):
        item = OrderItem(name, price, quantity)
        self._items.append(item)

    def get_items(self):
        return self._items

    def get_total(self):
        return sum(item.get_total() for item in self._items)

class OrderPresenter:
    def display_item_added(self, item):
        print(f"Added {item.get_quantity()} x {item.get_name()} to order")

    def display_order_summary(self, order):
        print("Order Summary:")
        for item in order.get_items():
            print(f"- {item.get_quantity()} x {item.get_name()} @ ${item.get_price():.2f}")
        print(f"Total: ${order.get_total():.2f}")

class OrderService:
    def __init__(self):
        self._order = OrderDomainAfter()
        self._presenter = OrderPresenter()

    def add_item(self, name, price, quantity):
        item = OrderItem(name, price, quantity)
        self._order.add_item(name, price, quantity)
        self._presenter.display_item_added(item)
        self._presenter.display_order_summary(self._order)

    def get_order(self):
        return self._order

# 72. Hierarchy Extraction (Extract Hierarchy)
#
# BEFORE: Single class handling multiple responsibilities
class ComputerExtractBefore:
    def __init__(self, computer_type, cpu, ram, storage):
        self._type = computer_type
        self._cpu = cpu
        self._ram = ram
        self._storage = storage

    def get_specs(self):
        specs = f"CPU: {self._cpu}\n"
        specs += f"RAM: {self._ram}GB\n"
        specs += f"Storage: {self._storage}GB\n"

        if self._type == 'desktop':
            specs += "Form Factor: Desktop\n"
        elif self._type == 'laptop':
            specs += "Form Factor: Laptop\n"
            specs += "Battery Life: 8 hours\n"
        elif self._type == 'server':
            specs += "Form Factor: Server Rack\n"
            specs += "Redundancy: RAID 10\n"

        return specs

# AFTER: Extract hierarchy
class ComputerExtractAfter(ABC):
    def __init__(self, cpu, ram, storage):
        self._cpu = cpu
        self._ram = ram
        self._storage = storage

    @abstractmethod
    def get_form_factor(self):
        pass

    @abstractmethod
    def get_special_features(self):
        pass

    def get_basic_specs(self):
        return f"CPU: {self._cpu}\nRAM: {self._ram}GB\nStorage: {self._storage}GB\n"

    def get_specs(self):
        return (self.get_basic_specs() +
                f"Form Factor: {self.get_form_factor()}\n" +
                self.get_special_features())

class DesktopComputer(ComputerExtractAfter):
    def get_form_factor(self):
        return "Desktop"

    def get_special_features(self):
        return "Expansion Slots: Multiple PCI\n"

class LaptopComputer(ComputerExtractAfter):
    def get_form_factor(self):
        return "Laptop"

    def get_special_features(self):
        return "Battery Life: 8 hours\nWeight: 2.5 lbs\n"

class ServerComputer(ComputerExtractAfter):
    def get_form_factor(self):
        return "Server Rack"

    def get_special_features(self):
        return "Redundancy: RAID 10\nHot Swap Drives: Yes\n"
