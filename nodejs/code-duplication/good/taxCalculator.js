class TaxCalculator {
    static TAX_RATES = {
        'CA': 0.0825,
        'NY': 0.04,
        'TX': 0.0625,
    };

    calculateTax(price, state) {
        const taxRate = TaxCalculator.TAX_RATES[state] || 0.0;
        return price * taxRate;
    }
}

class ShippingCalculator {
    static BASE_RATE = 5.0;
    static WEIGHT_RATE_PER_UNIT = 0.5;
    static DISTANCE_RATE_PER_UNIT = 0.1;

    calculateShipping(weight, distance) {
        const weightRate = weight * ShippingCalculator.WEIGHT_RATE_PER_UNIT;
        const distanceRate = distance * ShippingCalculator.DISTANCE_RATE_PER_UNIT;

        return ShippingCalculator.BASE_RATE + weightRate + distanceRate;
    }
}

module.exports = {
    TaxCalculator,
    ShippingCalculator
};
