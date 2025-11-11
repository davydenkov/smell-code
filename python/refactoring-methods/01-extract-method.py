# 1. Method Extraction (Extract Method)
#
# BEFORE: A method contains too much logic, making it hard to understand
class OrderProcessorBefore:
    def process_order(self, order):
        # Validate order
        if order['total'] <= 0:
            raise ValueError('Invalid order total')

        # Calculate tax
        tax = order['subtotal'] * 0.08

        # Calculate shipping
        shipping = 15.00 if order['weight'] > 10 else 5.00

        # Calculate total
        total = order['subtotal'] + tax + shipping

        # Save to database
        self.save_order(order, total)

        return total

    def save_order(self, order, total):
        # Database save logic
        pass

# AFTER: Extract methods to separate concerns
class OrderProcessorAfter:
    def process_order(self, order):
        self.validate_order(order)

        tax = self.calculate_tax(order)
        shipping = self.calculate_shipping(order)
        total = self.calculate_total(order, tax, shipping)

        self.save_order(order, total)

        return total

    def validate_order(self, order):
        if order['total'] <= 0:
            raise ValueError('Invalid order total')

    def calculate_tax(self, order):
        return order['subtotal'] * 0.08

    def calculate_shipping(self, order):
        return 15.00 if order['weight'] > 10 else 5.00

    def calculate_total(self, order, tax, shipping):
        return order['subtotal'] + tax + shipping

    def save_order(self, order, total):
        # Database save logic
        pass

# 2. Embedding a method (Inline Method)
#
# BEFORE: A method is too simple and adds no value
class UserBefore:
    def __init__(self, first_name, last_name):
        self.first_name = first_name
        self.last_name = last_name

    def get_full_name(self):
        return self.get_first_name() + ' ' + self.get_last_name()

    def get_first_name(self):
        return self.first_name

    def get_last_name(self):
        return self.last_name

# AFTER: Inline the simple method
class UserAfter:
    def __init__(self, first_name, last_name):
        self.first_name = first_name
        self.last_name = last_name

    def get_full_name(self):
        return self.first_name + ' ' + self.last_name

    # get_first_name() and get_last_name() methods removed
