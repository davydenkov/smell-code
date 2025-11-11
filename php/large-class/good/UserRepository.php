<?php

class UserRepository {
    private $db;

    public function __construct($db) {
        $this->db = $db;
    }

    public function create(string $email, string $name, string $hashedPassword): int {
        $stmt = $this->db->prepare("INSERT INTO users (email, name, password) VALUES (?, ?, ?)");
        $stmt->execute([$email, $name, $hashedPassword]);
        return $this->db->lastInsertId();
    }

    public function findByEmail(string $email): ?User {
        $stmt = $this->db->prepare("SELECT id, email, name, balance FROM users WHERE email = ?");
        $stmt->execute([$email]);
        $data = $stmt->fetch();

        if (!$data) {
            return null;
        }

        return new User($data['id'], $data['email'], $data['name'], $data['balance']);
    }

    public function findById(int $id): ?User {
        $stmt = $this->db->prepare("SELECT id, email, name, balance FROM users WHERE id = ?");
        $stmt->execute([$id]);
        $data = $stmt->fetch();

        if (!$data) {
            return null;
        }

        return new User($data['id'], $data['email'], $data['name'], $data['balance']);
    }

    public function update(User $user): void {
        $stmt = $this->db->prepare("UPDATE users SET name = ?, email = ? WHERE id = ?");
        $stmt->execute([$user->getName(), $user->getEmail(), $user->getId()]);
    }

    public function authenticate(string $email, string $password): ?User {
        $stmt = $this->db->prepare("SELECT id, email, name, password, balance FROM users WHERE email = ?");
        $stmt->execute([$email]);
        $data = $stmt->fetch();

        if ($data && password_verify($password, $data['password'])) {
            return new User($data['id'], $data['email'], $data['name'], $data['balance']);
        }

        return null;
    }
}
