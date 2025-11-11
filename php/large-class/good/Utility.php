<?php

class Utility {
    public static function validateEmail(string $email): bool {
        return filter_var($email, FILTER_VALIDATE_EMAIL);
    }

    public static function generateToken(): string {
        return bin2hex(random_bytes(32));
    }

    public static function hashPassword(string $password): string {
        return password_hash($password, PASSWORD_DEFAULT);
    }
}

class ActivityLogger {
    private $db;

    public function __construct($db) {
        $this->db = $db;
    }

    public function logActivity(int $userId, string $action): void {
        $stmt = $this->db->prepare("INSERT INTO activity_log (user_id, action, created_at) VALUES (?, ?, NOW())");
        $stmt->execute([$userId, $action]);
    }
}
