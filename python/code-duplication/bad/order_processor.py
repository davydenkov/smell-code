class OrderProcessor:
    def calculate_tax(self, price, state):
        tax_rate = 0.0

        if state == 'CA':
            tax_rate = 0.0825
        elif state == 'NY':
            tax_rate = 0.04
        elif state == 'TX':
            tax_rate = 0.0625

        return price * tax_rate

    def calculate_shipping(self, weight, distance):
        base_rate = 5.0
        weight_rate = weight * 0.5
        distance_rate = distance * 0.1

        return base_rate + weight_rate + distance_rate


class InvoiceGenerator:
    def calculate_tax(self, price, state):
        tax_rate = 0.0

        if state == 'CA':
            tax_rate = 0.0825
        elif state == 'NY':
            tax_rate = 0.04
        elif state == 'TX':
            tax_rate = 0.0625

        return price * tax_rate

    def calculate_shipping(self, weight, distance):
        base_rate = 5.0
        weight_rate = weight * 0.5
        distance_rate = distance * 0.1

        return base_rate + weight_rate + distance_rate


class QuoteGenerator:
    def calculate_tax(self, price, state):
        tax_rate = 0.0

        if state == 'CA':
            tax_rate = 0.0825
        elif state == 'NY':
            tax_rate = 0.04
        elif state == 'TX':
            tax_rate = 0.0625

        return price * tax_rate

    def calculate_shipping(self, weight, distance):
        base_rate = 5.0
        weight_rate = weight * 0.5
        distance_rate = distance * 0.1

        return base_rate + weight_rate + distance_rate
