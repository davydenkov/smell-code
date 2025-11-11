class TaxCalculator:
    TAX_RATES = {
        'CA': 0.0825,
        'NY': 0.04,
        'TX': 0.0625,
    }

    def calculate_tax(self, price: float, state: str) -> float:
        tax_rate = self.TAX_RATES.get(state, 0.0)
        return price * tax_rate


class ShippingCalculator:
    BASE_RATE = 5.0
    WEIGHT_RATE_PER_UNIT = 0.5
    DISTANCE_RATE_PER_UNIT = 0.1

    def calculate_shipping(self, weight: float, distance: float) -> float:
        weight_rate = weight * self.WEIGHT_RATE_PER_UNIT
        distance_rate = distance * self.DISTANCE_RATE_PER_UNIT

        return self.BASE_RATE + weight_rate + distance_rate
