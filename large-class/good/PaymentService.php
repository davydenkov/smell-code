<?php

class PaymentService {
    private string $stripeSecretKey;
    private string $paypalClientId;
    private string $paypalSecret;

    public function __construct(array $paymentConfig) {
        $this->stripeSecretKey = $paymentConfig['stripe_secret'];
        $this->paypalClientId = $paymentConfig['paypal_client_id'];
        $this->paypalSecret = $paymentConfig['paypal_secret'];
    }

    public function processStripePayment(float $amount, string $token): array {
        // Simulate Stripe payment processing
        if ($token && $amount > 0) {
            return ['success' => true, 'transaction_id' => 'stripe_' . uniqid()];
        }
        return ['success' => false, 'error' => 'Invalid payment data'];
    }

    public function processPayPalPayment(float $amount, string $paypalToken): array {
        // Simulate PayPal payment processing
        if ($paypalToken && $amount > 0) {
            return ['success' => true, 'transaction_id' => 'paypal_' . uniqid()];
        }
        return ['success' => false, 'error' => 'Invalid payment data'];
    }

    public function refundPayment(string $transactionId, float $amount): array {
        if (strpos($transactionId, 'stripe_') === 0) {
            return ['success' => true, 'refund_id' => 'refund_' . uniqid()];
        } elseif (strpos($transactionId, 'paypal_') === 0) {
            return ['success' => true, 'refund_id' => 'refund_' . uniqid()];
        }
        return ['success' => false, 'error' => 'Unknown transaction type'];
    }
}
