<?php

class UserValidator {
    public function validateRegistrationData(array $userData): void {
        $this->validateRequiredFields($userData);
        $this->validateEmailFormat($userData['email']);
        $this->validatePasswordStrength($userData['password']);
    }

    private function validateRequiredFields(array $userData): void {
        if (empty($userData['email'])) {
            throw new Exception('Email is required');
        }
        if (empty($userData['password'])) {
            throw new Exception('Password is required');
        }
    }

    private function validateEmailFormat(string $email): void {
        if (!filter_var($email, FILTER_VALIDATE_EMAIL)) {
            throw new Exception('Invalid email format');
        }
    }

    private function validatePasswordStrength(string $password): void {
        if (strlen($password) < 8) {
            throw new Exception('Password must be at least 8 characters');
        }
    }
}
