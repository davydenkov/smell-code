class OrderProcessor {
    calculateTax(price, state) {
        let taxRate = 0.0;

        if (state === 'CA') {
            taxRate = 0.0825;
        } else if (state === 'NY') {
            taxRate = 0.04;
        } else if (state === 'TX') {
            taxRate = 0.0625;
        }

        return price * taxRate;
    }

    calculateShipping(weight, distance) {
        const baseRate = 5.0;
        const weightRate = weight * 0.5;
        const distanceRate = distance * 0.1;

        return baseRate + weightRate + distanceRate;
    }
}

class InvoiceGenerator {
    calculateTax(price, state) {
        let taxRate = 0.0;

        if (state === 'CA') {
            taxRate = 0.0825;
        } else if (state === 'NY') {
            taxRate = 0.04;
        } else if (state === 'TX') {
            taxRate = 0.0625;
        }

        return price * taxRate;
    }

    calculateShipping(weight, distance) {
        const baseRate = 5.0;
        const weightRate = weight * 0.5;
        const distanceRate = distance * 0.1;

        return baseRate + weightRate + distanceRate;
    }
}

class QuoteGenerator {
    calculateTax(price, state) {
        let taxRate = 0.0;

        if (state === 'CA') {
            taxRate = 0.0825;
        } else if (state === 'NY') {
            taxRate = 0.04;
        } else if (state === 'TX') {
            taxRate = 0.0625;
        }

        return price * taxRate;
    }

    calculateShipping(weight, distance) {
        const baseRate = 5.0;
        const weightRate = weight * 0.5;
        const distanceRate = distance * 0.1;

        return baseRate + weightRate + distanceRate;
    }
}

module.exports = {
    OrderProcessor,
    InvoiceGenerator,
    QuoteGenerator
};
