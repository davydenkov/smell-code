/**
 * 3. Embedding a temporary variable (Inline Temp)
 *
 * BEFORE: Unnecessary temporary variable
 */
class PriceCalculatorBefore {
    getPrice() {
        const basePrice = this.quantity * this.itemPrice;
        if (basePrice > 1000) {
            return basePrice * 0.95;
        } else {
            return basePrice * 0.98;
        }
    }
}

/**
 * AFTER: Inline the temporary variable
 */
class PriceCalculatorAfter {
    getPrice() {
        if (this.quantity * this.itemPrice > 1000) {
            return this.quantity * this.itemPrice * 0.95;
        } else {
            return this.quantity * this.itemPrice * 0.98;
        }
    }
}

/**
 * 4. Replacing a temporary variable with a method call (Replace Temp with Query)
 *
 * BEFORE: Temporary variable used multiple times
 */
class OrderBefore {
    constructor() {
        this.quantity = 0;
        this.itemPrice = 0;
    }

    getPrice() {
        const basePrice = this.quantity * this.itemPrice;
        return basePrice - this.getDiscount(basePrice);
    }

    getDiscount(basePrice) {
        return Math.max(0, basePrice - 500) * 0.05;
    }
}

/**
 * AFTER: Replace temp with query
 */
class OrderAfter {
    constructor() {
        this.quantity = 0;
        this.itemPrice = 0;
    }

    getPrice() {
        return this.getBasePrice() - this.getDiscount();
    }

    getBasePrice() {
        return this.quantity * this.itemPrice;
    }

    getDiscount() {
        return Math.max(0, this.getBasePrice() - 500) * 0.05;
    }
}

/**
 * 5. Introducing an explanatory variable (Introduce Explaining Variable)
 *
 * BEFORE: Complex expression
 */
class PaymentCalculatorBefore {
    isEligibleForDiscount(customer) {
        return (customer.purchaseHistory.length > 5 &&
                customer.loyaltyPoints > 1000 &&
                customer.membershipType === 'premium');
    }
}

/**
 * AFTER: Introduce explaining variable
 */
class PaymentCalculatorAfter {
    isEligibleForDiscount(customer) {
        const hasLongPurchaseHistory = customer.purchaseHistory.length > 5;
        const hasHighLoyaltyPoints = customer.loyaltyPoints > 1000;
        const isPremiumMember = customer.membershipType === 'premium';

        return hasLongPurchaseHistory && hasHighLoyaltyPoints && isPremiumMember;
    }
}

/**
 * 6. Decomposing a complex expression (Extract Method)
 */
class InvoiceBefore {
    calculateTotal(invoice) {
        return invoice.subtotal +
               (invoice.subtotal * invoice.taxRate) +
               (invoice.subtotal * invoice.shippingRate);
    }
}

class InvoiceAfter {
    calculateTotal(invoice) {
        return invoice.subtotal + this.calculateTax(invoice) + this.calculateShipping(invoice);
    }

    calculateTax(invoice) {
        return invoice.subtotal * invoice.taxRate;
    }

    calculateShipping(invoice) {
        return invoice.subtotal * invoice.shippingRate;
    }
}

module.exports = {
    PriceCalculatorBefore,
    PriceCalculatorAfter,
    OrderBefore,
    OrderAfter,
    PaymentCalculatorBefore,
    PaymentCalculatorAfter,
    InvoiceBefore,
    InvoiceAfter
};
