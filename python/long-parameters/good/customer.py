class Customer:
    def __init__(
        self,
        id: int,
        name: str,
        email: str,
        phone: str = None,
        shipping_address=None,
        billing_address=None
    ):
        self.id = id
        self.name = name
        self.email = email
        self.phone = phone
        self.shipping_address = shipping_address or Address()
        self.billing_address = billing_address or Address()


class Address:
    def __init__(self, street: str = '', city: str = '', state: str = '', zip_code: str = ''):
        self.street = street
        self.city = city
        self.state = state
        self.zip_code = zip_code


class Product:
    def __init__(self, id: int, name: str, price: float):
        self.id = id
        self.name = name
        self.price = price


class OrderDetails:
    def __init__(
        self,
        product: Product,
        quantity: int,
        tax_rate: float = 0.0,
        discount_percent: float = 0.0,
        shipping_method: str = 'standard',
        shipping_cost: float = 0.0,
        payment_method: str = 'credit_card',
        notes: str = None
    ):
        self.product = product
        self.quantity = quantity
        self.tax_rate = tax_rate
        self.discount_percent = discount_percent
        self.shipping_method = shipping_method
        self.shipping_cost = shipping_cost
        self.payment_method = payment_method
        self.notes = notes


class OrderCalculator:
    def calculate_totals(self, order_details: OrderDetails):
        subtotal = order_details.product.price * order_details.quantity
        discount_amount = subtotal * (order_details.discount_percent / 100)
        taxable_amount = subtotal - discount_amount
        tax_amount = taxable_amount * (order_details.tax_rate / 100)
        total = taxable_amount + tax_amount + order_details.shipping_cost

        return {
            'subtotal': subtotal,
            'discount_amount': discount_amount,
            'tax_amount': tax_amount,
            'shipping_cost': order_details.shipping_cost,
            'total': total
        }
