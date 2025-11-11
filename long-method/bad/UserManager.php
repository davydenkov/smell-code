<?php

class UserManager {
    private $db;

    public function __construct($db) {
        $this->db = $db;
    }

    public function registerUser($userData) {
        // Validate input data
        if (empty($userData['email'])) {
            throw new Exception('Email is required');
        }
        if (empty($userData['password'])) {
            throw new Exception('Password is required');
        }
        if (strlen($userData['password']) < 8) {
            throw new Exception('Password must be at least 8 characters');
        }
        if (!filter_var($userData['email'], FILTER_VALIDATE_EMAIL)) {
            throw new Exception('Invalid email format');
        }

        // Check if user already exists
        $stmt = $this->db->prepare("SELECT id FROM users WHERE email = ?");
        $stmt->execute([$userData['email']]);
        if ($stmt->fetch()) {
            throw new Exception('User already exists');
        }

        // Hash password
        $hashedPassword = password_hash($userData['password'], PASSWORD_DEFAULT);

        // Generate verification token
        $verificationToken = bin2hex(random_bytes(32));

        // Insert user into database
        $stmt = $this->db->prepare("INSERT INTO users (email, password, first_name, last_name, verification_token, created_at) VALUES (?, ?, ?, ?, ?, NOW())");
        $stmt->execute([
            $userData['email'],
            $hashedPassword,
            $userData['firstName'],
            $userData['lastName'],
            $verificationToken
        ]);

        $userId = $this->db->lastInsertId();

        // Create user profile
        $stmt = $this->db->prepare("INSERT INTO user_profiles (user_id, phone, address, city, state, zip_code) VALUES (?, ?, ?, ?, ?, ?)");
        $stmt->execute([
            $userId,
            $userData['phone'] ?? null,
            $userData['address'] ?? null,
            $userData['city'] ?? null,
            $userData['state'] ?? null,
            $userData['zipCode'] ?? null
        ]);

        // Send verification email
        $subject = 'Please verify your email address';
        $message = "Hello {$userData['firstName']},\n\n";
        $message .= "Thank you for registering. Please click the link below to verify your email:\n\n";
        $message .= "http://example.com/verify?token={$verificationToken}\n\n";
        $message .= "Best regards,\nThe Team";

        $headers = 'From: noreply@example.com' . "\r\n" .
                   'Reply-To: noreply@example.com' . "\r\n" .
                   'X-Mailer: PHP/' . phpversion();

        mail($userData['email'], $subject, $message, $headers);

        // Log registration
        $logMessage = "User registered: {$userData['email']} at " . date('Y-m-d H:i:s');
        error_log($logMessage);

        // Create default settings
        $stmt = $this->db->prepare("INSERT INTO user_settings (user_id, theme, notifications_enabled) VALUES (?, 'light', true)");
        $stmt->execute([$userId]);

        // Send welcome notification
        $stmt = $this->db->prepare("INSERT INTO notifications (user_id, type, message, created_at) VALUES (?, 'welcome', 'Welcome to our platform!', NOW())");
        $stmt->execute([$userId]);

        return $userId;
    }
}
