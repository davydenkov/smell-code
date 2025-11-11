<?php

class NotificationService {
    private $db;

    public function __construct($db) {
        $this->db = $db;
    }

    public function sendWelcomeNotification(int $userId): void {
        $stmt = $this->db->prepare(
            "INSERT INTO notifications (user_id, type, message, created_at)
             VALUES (?, 'welcome', 'Welcome to our platform!', NOW())"
        );
        $stmt->execute([$userId]);
    }
}

class Logger {
    public static function logRegistration(string $email): void {
        $logMessage = "User registered: {$email} at " . date('Y-m-d H:i:s');
        error_log($logMessage);
    }
}
