class OrderProcessor:
    def __init__(self):
        self.tax_calculator = TaxCalculator()
        self.shipping_calculator = ShippingCalculator()

    def calculate_tax(self, price: float, state: str) -> float:
        return self.tax_calculator.calculate_tax(price, state)

    def calculate_shipping(self, weight: float, distance: float) -> float:
        return self.shipping_calculator.calculate_shipping(weight, distance)


class InvoiceGenerator:
    def __init__(self):
        self.tax_calculator = TaxCalculator()
        self.shipping_calculator = ShippingCalculator()

    def calculate_tax(self, price: float, state: str) -> float:
        return self.tax_calculator.calculate_tax(price, state)

    def calculate_shipping(self, weight: float, distance: float) -> float:
        return self.shipping_calculator.calculate_shipping(weight, distance)


class QuoteGenerator:
    def __init__(self):
        self.tax_calculator = TaxCalculator()
        self.shipping_calculator = ShippingCalculator()

    def calculate_tax(self, price: float, state: str) -> float:
        return self.tax_calculator.calculate_tax(price, state)

    def calculate_shipping(self, weight: float, distance: float) -> float:
        return self.shipping_calculator.calculate_shipping(weight, distance)
