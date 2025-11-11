<?php

class User {
    private int $id;
    private string $email;
    private string $name;
    private float $balance;

    public function __construct(int $id, string $email, string $name, float $balance = 0.0) {
        $this->id = $id;
        $this->email = $email;
        $this->name = $name;
        $this->balance = $balance;
    }

    public function getId(): int {
        return $this->id;
    }

    public function getEmail(): string {
        return $this->email;
    }

    public function getName(): string {
        return $this->name;
    }

    public function getBalance(): float {
        return $this->balance;
    }

    public function updateProfile(string $name, string $email): void {
        $this->name = $name;
        $this->email = $email;
    }
}
