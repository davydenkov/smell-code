<?php

class OrderService {
    private OrderCalculator $calculator;

    public function __construct() {
        $this->calculator = new OrderCalculator();
    }

    public function createOrder(Customer $customer, OrderDetails $orderDetails): array {
        $totals = $this->calculator->calculateTotals($orderDetails);

        $orderData = [
            'customer_id' => $customer->id,
            'customer_name' => $customer->name,
            'customer_email' => $customer->email,
            'customer_phone' => $customer->phone,
            'customer_address' => $customer->shippingAddress->street,
            'customer_city' => $customer->shippingAddress->city,
            'customer_state' => $customer->shippingAddress->state,
            'customer_zip' => $customer->shippingAddress->zipCode,
            'product_id' => $orderDetails->product->id,
            'product_name' => $orderDetails->product->name,
            'product_price' => $orderDetails->product->price,
            'quantity' => $orderDetails->quantity,
            'subtotal' => $totals['subtotal'],
            'discount_percent' => $orderDetails->discountPercent,
            'discount_amount' => $totals['discount_amount'],
            'tax_rate' => $orderDetails->taxRate,
            'tax_amount' => $totals['tax_amount'],
            'shipping_method' => $orderDetails->shippingMethod,
            'shipping_cost' => $totals['shipping_cost'],
            'payment_method' => $orderDetails->paymentMethod,
            'billing_address' => $customer->billingAddress->street,
            'billing_city' => $customer->billingAddress->city,
            'billing_state' => $customer->billingAddress->state,
            'billing_zip' => $customer->billingAddress->zipCode,
            'total' => $totals['total'],
            'notes' => $orderDetails->notes,
            'created_at' => date('Y-m-d H:i:s')
        ];

        // In a real application, this would save to database
        return $orderData;
    }

    public function updateOrder(int $orderId, Customer $customer, OrderDetails $orderDetails): array {
        $totals = $this->calculator->calculateTotals($orderDetails);

        $orderData = [
            'id' => $orderId,
            'customer_id' => $customer->id,
            'customer_name' => $customer->name,
            'customer_email' => $customer->email,
            'customer_phone' => $customer->phone,
            'customer_address' => $customer->shippingAddress->street,
            'customer_city' => $customer->shippingAddress->city,
            'customer_state' => $customer->shippingAddress->state,
            'customer_zip' => $customer->shippingAddress->zipCode,
            'product_id' => $orderDetails->product->id,
            'product_name' => $orderDetails->product->name,
            'product_price' => $orderDetails->product->price,
            'quantity' => $orderDetails->quantity,
            'subtotal' => $totals['subtotal'],
            'discount_percent' => $orderDetails->discountPercent,
            'discount_amount' => $totals['discount_amount'],
            'tax_rate' => $orderDetails->taxRate,
            'tax_amount' => $totals['tax_amount'],
            'shipping_method' => $orderDetails->shippingMethod,
            'shipping_cost' => $totals['shipping_cost'],
            'payment_method' => $orderDetails->paymentMethod,
            'billing_address' => $customer->billingAddress->street,
            'billing_city' => $customer->billingAddress->city,
            'billing_state' => $customer->billingAddress->state,
            'billing_zip' => $customer->billingAddress->zipCode,
            'total' => $totals['total'],
            'notes' => $orderDetails->notes,
            'updated_at' => date('Y-m-d H:i:s')
        ];

        return $orderData;
    }
}
