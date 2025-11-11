const { TaxCalculator, ShippingCalculator } = require('./taxCalculator');

class OrderProcessor {
    constructor() {
        this.taxCalculator = new TaxCalculator();
        this.shippingCalculator = new ShippingCalculator();
    }

    calculateTax(price, state) {
        return this.taxCalculator.calculateTax(price, state);
    }

    calculateShipping(weight, distance) {
        return this.shippingCalculator.calculateShipping(weight, distance);
    }
}

class InvoiceGenerator {
    constructor() {
        this.taxCalculator = new TaxCalculator();
        this.shippingCalculator = new ShippingCalculator();
    }

    calculateTax(price, state) {
        return this.taxCalculator.calculateTax(price, state);
    }

    calculateShipping(weight, distance) {
        return this.shippingCalculator.calculateShipping(weight, distance);
    }
}

class QuoteGenerator {
    constructor() {
        this.taxCalculator = new TaxCalculator();
        this.shippingCalculator = new ShippingCalculator();
    }

    calculateTax(price, state) {
        return this.taxCalculator.calculateTax(price, state);
    }

    calculateShipping(weight, distance) {
        return this.shippingCalculator.calculateShipping(weight, distance);
    }
}

module.exports = {
    OrderProcessor,
    InvoiceGenerator,
    QuoteGenerator
};
