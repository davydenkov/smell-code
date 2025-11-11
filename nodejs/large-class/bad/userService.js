// BAD: Large class that does too many things (violates Single Responsibility Principle)
class UserService {
    constructor(db, emailConfig, paymentConfig) {
        this.db = db;
        this.emailConfig = emailConfig;
        this.paymentConfig = paymentConfig;

        // User properties
        this.userId = null;
        this.userEmail = '';
        this.userName = '';
        this.userBalance = 0;

        // Email properties
        this.smtpHost = emailConfig.host;
        this.smtpPort = emailConfig.port;
        this.smtpUsername = emailConfig.username;
        this.smtpPassword = emailConfig.password;

        // Payment properties
        this.stripeSecretKey = paymentConfig.stripeSecretKey;
        this.paypalClientId = paymentConfig.paypalClientId;
        this.paypalSecret = paymentConfig.paypalSecret;
    }

    // User management methods
    async createUser(userData) {
        // Validate user data
        this.validateUserData(userData);

        // Hash password
        const bcrypt = require('bcrypt');
        userData.password = await bcrypt.hash(userData.password, 10);

        // Insert user
        const result = await this.db.query(
            'INSERT INTO users (email, password, name) VALUES ($1, $2, $3) RETURNING id',
            [userData.email, userData.password, userData.name]
        );

        return result.rows[0].id;
    }

    validateUserData(userData) {
        if (!userData.email || !userData.password || !userData.name) {
            throw new Error('Missing required fields');
        }
    }

    async getUser(userId) {
        const result = await this.db.query('SELECT * FROM users WHERE id = $1', [userId]);
        return result.rows[0];
    }

    // Email methods
    async sendWelcomeEmail(userId) {
        const user = await this.getUser(userId);
        const nodemailer = require('nodemailer');

        const transporter = nodemailer.createTransporter({
            host: this.smtpHost,
            port: this.smtpPort,
            auth: {
                user: this.smtpUsername,
                pass: this.smtpPassword
            }
        });

        await transporter.sendMail({
            from: this.smtpUsername,
            to: user.email,
            subject: 'Welcome!',
            text: `Welcome ${user.name}!`
        });
    }

    // Payment methods
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

    // Reporting methods
    async generateUserReport(userId) {
        const user = await this.getUser(userId);
        const payments = await this.db.query('SELECT * FROM payments WHERE user_id = $1', [userId]);

        return {
            user: user,
            payments: payments.rows,
            totalSpent: payments.rows.reduce((sum, payment) => sum + payment.amount, 0),
            reportGenerated: new Date()
        };
    }

    // Utility methods
    formatCurrency(amount) {
        return `$${amount.toFixed(2)}`;
    }

    calculateDiscount(amount, percentage) {
        return amount * (1 - percentage / 100);
    }

    // Many more methods would be here in a real large class...
}

module.exports = UserService;
