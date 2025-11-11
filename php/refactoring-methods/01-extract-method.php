<?php

/**
 * 1. Method Extraction (Extract Method)
 *
 * BEFORE: A method contains too much logic, making it hard to understand
 */
class OrderProcessorBefore {
    public function processOrder($order) {
        // Validate order
        if ($order['total'] <= 0) {
            throw new Exception('Invalid order total');
        }

        // Calculate tax
        $tax = $order['subtotal'] * 0.08;

        // Calculate shipping
        $shipping = $order['weight'] > 10 ? 15.00 : 5.00;

        // Calculate total
        $total = $order['subtotal'] + $tax + $shipping;

        // Save to database
        $this->saveOrder($order, $total);

        return $total;
    }

    private function saveOrder($order, $total) {
        // Database save logic
    }
}

/**
 * AFTER: Extract methods to separate concerns
 */
class OrderProcessorAfter {
    public function processOrder($order) {
        $this->validateOrder($order);

        $tax = $this->calculateTax($order);
        $shipping = $this->calculateShipping($order);
        $total = $this->calculateTotal($order, $tax, $shipping);

        $this->saveOrder($order, $total);

        return $total;
    }

    private function validateOrder($order) {
        if ($order['total'] <= 0) {
            throw new Exception('Invalid order total');
        }
    }

    private function calculateTax($order) {
        return $order['subtotal'] * 0.08;
    }

    private function calculateShipping($order) {
        return $order['weight'] > 10 ? 15.00 : 5.00;
    }

    private function calculateTotal($order, $tax, $shipping) {
        return $order['subtotal'] + $tax + $shipping;
    }

    private function saveOrder($order, $total) {
        // Database save logic
    }
}

/**
 * 2. Embedding a method (Inline Method)
 *
 * BEFORE: A method is too simple and adds no value
 */
class UserBefore {
    public function getFullName() {
        return $this->getFirstName() . ' ' . $this->getLastName();
    }

    public function getFirstName() {
        return $this->firstName;
    }

    public function getLastName() {
        return $this->lastName;
    }
}

/**
 * AFTER: Inline the simple method
 */
class UserAfter {
    public function getFullName() {
        return $this->firstName . ' ' . $this->lastName;
    }

    // getFirstName() and getLastName() methods removed
}
