<?php

class OrderCalculator {
    public function calculateTotals(OrderDetails $orderDetails): array {
        $subtotal = $orderDetails->product->price * $orderDetails->quantity;
        $discountAmount = $subtotal * ($orderDetails->discountPercent / 100);
        $taxableAmount = $subtotal - $discountAmount;
        $taxAmount = $taxableAmount * ($orderDetails->taxRate / 100);
        $total = $taxableAmount + $taxAmount + $orderDetails->shippingCost;

        return [
            'subtotal' => $subtotal,
            'discount_amount' => $discountAmount,
            'tax_amount' => $taxAmount,
            'shipping_cost' => $orderDetails->shippingCost,
            'total' => $total
        ];
    }
}
