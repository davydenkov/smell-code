<?php

class OrderService {
    public function createOrder(
        $customerId,
        $customerName,
        $customerEmail,
        $customerPhone,
        $customerAddress,
        $customerCity,
        $customerState,
        $customerZipCode,
        $productId,
        $productName,
        $productPrice,
        $quantity,
        $taxRate,
        $discountPercent,
        $shippingMethod,
        $shippingCost,
        $paymentMethod,
        $billingAddress,
        $billingCity,
        $billingState,
        $billingZipCode,
        $notes
    ) {
        // Calculate totals
        $subtotal = $productPrice * $quantity;
        $discountAmount = $subtotal * ($discountPercent / 100);
        $taxableAmount = $subtotal - $discountAmount;
        $taxAmount = $taxableAmount * ($taxRate / 100);
        $total = $taxableAmount + $taxAmount + $shippingCost;

        // Create order record
        $orderData = [
            'customer_id' => $customerId,
            'customer_name' => $customerName,
            'customer_email' => $customerEmail,
            'customer_phone' => $customerPhone,
            'customer_address' => $customerAddress,
            'customer_city' => $customerCity,
            'customer_state' => $customerState,
            'customer_zip' => $customerZipCode,
            'product_id' => $productId,
            'product_name' => $productName,
            'product_price' => $productPrice,
            'quantity' => $quantity,
            'subtotal' => $subtotal,
            'discount_percent' => $discountPercent,
            'discount_amount' => $discountAmount,
            'tax_rate' => $taxRate,
            'tax_amount' => $taxAmount,
            'shipping_method' => $shippingMethod,
            'shipping_cost' => $shippingCost,
            'payment_method' => $paymentMethod,
            'billing_address' => $billingAddress,
            'billing_city' => $billingCity,
            'billing_state' => $billingState,
            'billing_zip' => $billingZipCode,
            'total' => $total,
            'notes' => $notes,
            'created_at' => date('Y-m-d H:i:s')
        ];

        // In a real application, this would save to database
        return $orderData;
    }

    public function updateOrder(
        $orderId,
        $customerId,
        $customerName,
        $customerEmail,
        $customerPhone,
        $customerAddress,
        $customerCity,
        $customerState,
        $customerZipCode,
        $productId,
        $productName,
        $productPrice,
        $quantity,
        $taxRate,
        $discountPercent,
        $shippingMethod,
        $shippingCost,
        $paymentMethod,
        $billingAddress,
        $billingCity,
        $billingState,
        $billingZipCode,
        $notes
    ) {
        // Similar logic but for updating
        $subtotal = $productPrice * $quantity;
        $discountAmount = $subtotal * ($discountPercent / 100);
        $taxableAmount = $subtotal - $discountAmount;
        $taxAmount = $taxableAmount * ($taxRate / 100);
        $total = $taxableAmount + $taxAmount + $shippingCost;

        $orderData = [
            'id' => $orderId,
            'customer_id' => $customerId,
            'customer_name' => $customerName,
            'customer_email' => $customerEmail,
            'customer_phone' => $customerPhone,
            'customer_address' => $customerAddress,
            'customer_city' => $customerCity,
            'customer_state' => $customerState,
            'customer_zip' => $customerZipCode,
            'product_id' => $productId,
            'product_name' => $productName,
            'product_price' => $productPrice,
            'quantity' => $quantity,
            'subtotal' => $subtotal,
            'discount_percent' => $discountPercent,
            'discount_amount' => $discountAmount,
            'tax_rate' => $taxRate,
            'tax_amount' => $taxAmount,
            'shipping_method' => $shippingMethod,
            'shipping_cost' => $shippingCost,
            'payment_method' => $paymentMethod,
            'billing_address' => $billingAddress,
            'billing_city' => $billingCity,
            'billing_state' => $billingState,
            'billing_zip' => $billingZipCode,
            'total' => $total,
            'notes' => $notes,
            'updated_at' => date('Y-m-d H:i:s')
        ];

        return $orderData;
    }
}
