<?php

class UserRepository {
    private $db;

    public function __construct($db) {
        $this->db = $db;
    }

    public function userExists(string $email): bool {
        $stmt = $this->db->prepare("SELECT id FROM users WHERE email = ?");
        $stmt->execute([$email]);
        return $stmt->fetch() !== false;
    }

    public function createUser(array $userData): int {
        $stmt = $this->db->prepare(
            "INSERT INTO users (email, password, first_name, last_name, verification_token, created_at)
             VALUES (?, ?, ?, ?, ?, NOW())"
        );
        $stmt->execute([
            $userData['email'],
            $userData['password'],
            $userData['firstName'],
            $userData['lastName'],
            $userData['verificationToken']
        ]);

        return $this->db->lastInsertId();
    }

    public function createUserProfile(int $userId, array $profileData): void {
        $stmt = $this->db->prepare(
            "INSERT INTO user_profiles (user_id, phone, address, city, state, zip_code)
             VALUES (?, ?, ?, ?, ?, ?)"
        );
        $stmt->execute([
            $userId,
            $profileData['phone'] ?? null,
            $profileData['address'] ?? null,
            $profileData['city'] ?? null,
            $profileData['state'] ?? null,
            $profileData['zipCode'] ?? null
        ]);
    }

    public function createUserSettings(int $userId): void {
        $stmt = $this->db->prepare(
            "INSERT INTO user_settings (user_id, theme, notifications_enabled)
             VALUES (?, 'light', true)"
        );
        $stmt->execute([$userId]);
    }
}
