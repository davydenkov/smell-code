<?php

class UserService {
    private $db;
    private $emailConfig;
    private $paymentConfig;

    // User properties
    private $userId;
    private $userEmail;
    private $userName;
    private $userBalance;

    // Email properties
    private $smtpHost;
    private $smtpPort;
    private $smtpUsername;
    private $smtpPassword;

    // Payment properties
    private $stripeSecretKey;
    private $paypalClientId;
    private $paypalSecret;

    public function __construct($db, $emailConfig, $paymentConfig) {
        $this->db = $db;
        $this->emailConfig = $emailConfig;
        $this->paymentConfig = $paymentConfig;

        $this->smtpHost = $emailConfig['host'];
        $this->smtpPort = $emailConfig['port'];
        $this->smtpUsername = $emailConfig['username'];
        $this->smtpPassword = $emailConfig['password'];

        $this->stripeSecretKey = $paymentConfig['stripe_secret'];
        $this->paypalClientId = $paymentConfig['paypal_client_id'];
        $this->paypalSecret = $paymentConfig['paypal_secret'];
    }

    // User management methods
    public function createUser($email, $name, $password) {
        $hashedPassword = password_hash($password, PASSWORD_DEFAULT);
        $stmt = $this->db->prepare("INSERT INTO users (email, name, password) VALUES (?, ?, ?)");
        $stmt->execute([$email, $name, $hashedPassword]);
        return $this->db->lastInsertId();
    }

    public function authenticateUser($email, $password) {
        $stmt = $this->db->prepare("SELECT id, password FROM users WHERE email = ?");
        $stmt->execute([$email]);
        $user = $stmt->fetch();

        if ($user && password_verify($password, $user['password'])) {
            $this->userId = $user['id'];
            $this->userEmail = $email;
            return $user['id'];
        }
        return false;
    }

    public function updateUserProfile($userId, $name, $email) {
        $stmt = $this->db->prepare("UPDATE users SET name = ?, email = ? WHERE id = ?");
        $stmt->execute([$name, $email, $userId]);
    }

    public function getUserBalance($userId) {
        $stmt = $this->db->prepare("SELECT balance FROM users WHERE id = ?");
        $stmt->execute([$userId]);
        $result = $stmt->fetch();
        return $result['balance'] ?? 0;
    }

    // Email methods
    public function sendWelcomeEmail($email, $name) {
        $subject = 'Welcome to our platform!';
        $message = "Hello {$name},\n\nWelcome to our platform!";

        $headers = 'From: ' . $this->smtpUsername . "\r\n";
        $headers .= 'Reply-To: ' . $this->smtpUsername . "\r\n";

        return mail($email, $subject, $message, $headers);
    }

    public function sendPasswordResetEmail($email, $resetToken) {
        $subject = 'Password Reset';
        $message = "Click here to reset your password: http://example.com/reset?token={$resetToken}";

        $headers = 'From: ' . $this->smtpUsername . "\r\n";

        return mail($email, $subject, $message, $headers);
    }

    public function sendNotificationEmail($email, $subject, $message) {
        $headers = 'From: ' . $this->smtpUsername . "\r\n";

        return mail($email, $subject, $message, $headers);
    }

    // Payment methods
    public function processStripePayment($amount, $token) {
        // Simulate Stripe payment processing
        // In real code, this would use Stripe SDK
        if ($token && $amount > 0) {
            // Process payment logic here
            return ['success' => true, 'transaction_id' => 'stripe_' . uniqid()];
        }
        return ['success' => false, 'error' => 'Invalid payment data'];
    }

    public function processPayPalPayment($amount, $paypalToken) {
        // Simulate PayPal payment processing
        if ($paypalToken && $amount > 0) {
            // Process payment logic here
            return ['success' => true, 'transaction_id' => 'paypal_' . uniqid()];
        }
        return ['success' => false, 'error' => 'Invalid payment data'];
    }

    public function refundPayment($transactionId, $amount) {
        // Simulate refund logic
        if (strpos($transactionId, 'stripe_') === 0) {
            // Stripe refund
            return ['success' => true, 'refund_id' => 'refund_' . uniqid()];
        } elseif (strpos($transactionId, 'paypal_') === 0) {
            // PayPal refund
            return ['success' => true, 'refund_id' => 'refund_' . uniqid()];
        }
        return ['success' => false, 'error' => 'Unknown transaction type'];
    }

    // Reporting methods
    public function generateUserReport($userId) {
        $stmt = $this->db->prepare("
            SELECT u.name, u.email, u.created_at,
                   COUNT(o.id) as order_count,
                   SUM(o.total) as total_spent
            FROM users u
            LEFT JOIN orders o ON u.id = o.user_id
            WHERE u.id = ?
            GROUP BY u.id
        ");
        $stmt->execute([$userId]);
        return $stmt->fetch();
    }

    public function generateSalesReport($startDate, $endDate) {
        $stmt = $this->db->prepare("
            SELECT DATE(created_at) as date,
                   COUNT(*) as order_count,
                   SUM(total) as total_sales
            FROM orders
            WHERE created_at BETWEEN ? AND ?
            GROUP BY DATE(created_at)
        ");
        $stmt->execute([$startDate, $endDate]);
        return $stmt->fetchAll();
    }

    // Utility methods
    public function validateEmail($email) {
        return filter_var($email, FILTER_VALIDATE_EMAIL);
    }

    public function generateToken() {
        return bin2hex(random_bytes(32));
    }

    public function hashPassword($password) {
        return password_hash($password, PASSWORD_DEFAULT);
    }

    public function logActivity($userId, $action) {
        $stmt = $this->db->prepare("INSERT INTO activity_log (user_id, action, created_at) VALUES (?, ?, NOW())");
        $stmt->execute([$userId, $action]);
    }
}
