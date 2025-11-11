// BAD: Too many parameters (long parameter list smell)
class OrderService {
    createOrder(
        customerId,
        customerName,
        customerEmail,
        customerPhone,
        customerAddress,
        customerCity,
        customerState,
        customerZipCode,
        productId,
        productName,
        productPrice,
        quantity,
        taxRate,
        discountPercent,
        shippingMethod,
        shippingCost,
        paymentMethod,
        billingAddress,
        billingCity,
        billingState,
        billingZipCode,
        notes
    ) {
        // Calculate totals
        const subtotal = productPrice * quantity;
        const discountAmount = subtotal * (discountPercent / 100);
        const taxableAmount = subtotal - discountAmount;
        const taxAmount = taxableAmount * (taxRate / 100);
        const total = taxableAmount + taxAmount + shippingCost;

        // Create order record
        const orderData = {
            customer_id: customerId,
            customer_name: customerName,
            customer_email: customerEmail,
            customer_phone: customerPhone,
            customer_address: customerAddress,
            customer_city: customerCity,
            customer_state: customerState,
            customer_zip: customerZipCode,
            product_id: productId,
            product_name: productName,
            product_price: productPrice,
            quantity: quantity,
            subtotal: subtotal,
            discount_percent: discountPercent,
            discount_amount: discountAmount,
            tax_rate: taxRate,
            tax_amount: taxAmount,
            shipping_method: shippingMethod,
            shipping_cost: shippingCost,
            payment_method: paymentMethod,
            billing_address: billingAddress,
            billing_city: billingCity,
            billing_state: billingState,
            billing_zip: billingZipCode,
            total: total,
            notes: notes,
            created_at: new Date().toISOString()
        };

        // In a real application, this would save to database
        return orderData;
    }

    updateOrder(
        orderId,
        customerId,
        customerName,
        customerEmail,
        customerPhone,
        customerAddress,
        customerCity,
        customerState,
        customerZipCode,
        productId,
        productName,
        productPrice,
        quantity,
        taxRate,
        discountPercent,
        shippingMethod,
        shippingCost,
        paymentMethod,
        billingAddress,
        billingCity,
        billingState,
        billingZipCode,
        notes
    ) {
        // Similar logic but for updating
        const subtotal = productPrice * quantity;
        const discountAmount = subtotal * (discountPercent / 100);
        const taxableAmount = subtotal - discountAmount;
        const taxAmount = taxableAmount * (taxRate / 100);
        const total = taxableAmount + taxAmount + shippingCost;

        const orderData = {
            id: orderId,
            customer_id: customerId,
            customer_name: customerName,
            customer_email: customerEmail,
            customer_phone: customerPhone,
            customer_address: customerAddress,
            customer_city: customerCity,
            customer_state: customerState,
            customer_zip: customerZipCode,
            product_id: productId,
            product_name: productName,
            product_price: productPrice,
            quantity: quantity,
            subtotal: subtotal,
            discount_percent: discountPercent,
            discount_amount: discountAmount,
            tax_rate: taxRate,
            tax_amount: taxAmount,
            shipping_method: shippingMethod,
            shipping_cost: shippingCost,
            payment_method: paymentMethod,
            billing_address: billingAddress,
            billing_city: billingCity,
            billing_state: billingState,
            billing_zip: billingZipCode,
            total: total,
            notes: notes,
            updated_at: new Date().toISOString()
        };

        return orderData;
    }
}

module.exports = OrderService;
