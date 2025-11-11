class PaymentService {
    constructor(paymentConfig) {
        this.stripeSecretKey = paymentConfig.stripeSecretKey;
        this.paypalClientId = paymentConfig.paypalClientId;
        this.paypalSecret = paymentConfig.paypalSecret;
    }

    async processPayment(userId, amount, method) {
        if (method === 'stripe') {
            return this.processStripePayment(userId, amount);
        } else if (method === 'paypal') {
            return this.processPayPalPayment(userId, amount);
        }
        throw new Error('Unsupported payment method');
    }

    async processStripePayment(userId, amount) {
        // Stripe payment logic using this.stripeSecretKey
        console.log(`Processing $${amount} payment for user ${userId} via Stripe`);
        return { success: true, transactionId: 'stripe_txn_' + Date.now() };
    }

    async processPayPalPayment(userId, amount) {
        // PayPal payment logic using this.paypalClientId and this.paypalSecret
        console.log(`Processing $${amount} payment for user ${userId} via PayPal`);
        return { success: true, transactionId: 'paypal_txn_' + Date.now() };
    }
}

module.exports = PaymentService;
