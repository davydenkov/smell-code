/**
 * 1. Method Extraction (Extract Method)
 *
 * BEFORE: A method contains too much logic, making it hard to understand
 */
class OrderProcessorBefore {
    processOrder(order) {
        // Validate order
        if (order.total <= 0) {
            throw new Error('Invalid order total');
        }

        // Calculate tax
        const tax = order.subtotal * 0.08;

        // Calculate shipping
        const shipping = order.weight > 10 ? 15.00 : 5.00;

        // Calculate total
        const total = order.subtotal + tax + shipping;

        // Save to database
        this.saveOrder(order, total);

        return total;
    }

    saveOrder(order, total) {
        // Database save logic
    }
}

/**
 * AFTER: Extract methods to separate concerns
 */
class OrderProcessorAfter {
    processOrder(order) {
        this.validateOrder(order);

        const tax = this.calculateTax(order);
        const shipping = this.calculateShipping(order);
        const total = this.calculateTotal(order, tax, shipping);

        this.saveOrder(order, total);

        return total;
    }

    validateOrder(order) {
        if (order.total <= 0) {
            throw new Error('Invalid order total');
        }
    }

    calculateTax(order) {
        return order.subtotal * 0.08;
    }

    calculateShipping(order) {
        return order.weight > 10 ? 15.00 : 5.00;
    }

    calculateTotal(order, tax, shipping) {
        return order.subtotal + tax + shipping;
    }

    saveOrder(order, total) {
        // Database save logic
    }
}

/**
 * 2. Embedding a method (Inline Method)
 *
 * BEFORE: A method is too simple and adds no value
 */
class UserBefore {
    getFullName() {
        return this.getFirstName() + ' ' + this.getLastName();
    }

    getFirstName() {
        return this.firstName;
    }

    getLastName() {
        return this.lastName;
    }
}

/**
 * AFTER: Inline the simple method
 */
class UserAfter {
    getFullName() {
        return this.firstName + ' ' + this.lastName;
    }

    // getFirstName() and getLastName() methods removed
}

module.exports = {
    OrderProcessorBefore,
    OrderProcessorAfter,
    UserBefore,
    UserAfter
};
