class OrderCalculator {
    calculateTotals(orderDetails) {
        const subtotal = orderDetails.product.price * orderDetails.quantity;
        const discountAmount = subtotal * (orderDetails.discountPercent / 100);
        const taxableAmount = subtotal - discountAmount;
        const taxAmount = taxableAmount * (orderDetails.taxRate / 100);
        const total = taxableAmount + taxAmount + orderDetails.shippingCost;

        return {
            subtotal: subtotal,
            discount_amount: discountAmount,
            tax_amount: taxAmount,
            shipping_cost: orderDetails.shippingCost,
            total: total
        };
    }
}

module.exports = OrderCalculator;
