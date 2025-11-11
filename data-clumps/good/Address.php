<?php

class Address {
    private string $street;
    private string $city;
    private string $state;
    private string $zipCode;

    public function __construct(string $street, string $city, string $state, string $zipCode) {
        $this->street = $street;
        $this->city = $city;
        $this->state = $state;
        $this->zipCode = $zipCode;
    }

    public function getStreet(): string {
        return $this->street;
    }

    public function getCity(): string {
        return $this->city;
    }

    public function getState(): string {
        return $this->state;
    }

    public function getZipCode(): string {
        return $this->zipCode;
    }

    public function isValid(): bool {
        if (empty($this->street) || empty($this->city) || empty($this->state) || empty($this->zipCode)) {
            return false;
        }

        if (strlen($this->zipCode) !== 5) {
            return false;
        }

        return true;
    }

    public function toString(): string {
        return $this->street . ', ' . $this->city . ', ' . $this->state . ' ' . $this->zipCode;
    }

    public function toLabelFormat(): string {
        return $this->street . "\n" . $this->city . ', ' . $this->state . ' ' . $this->zipCode;
    }
}

class Person {
    private string $firstName;
    private string $lastName;
    private string $email;
    private ?string $phone;
    private ?string $dateOfBirth;

    public function __construct(
        string $firstName,
        string $lastName,
        string $email,
        ?string $phone = null,
        ?string $dateOfBirth = null
    ) {
        $this->firstName = $firstName;
        $this->lastName = $lastName;
        $this->email = $email;
        $this->phone = $phone;
        $this->dateOfBirth = $dateOfBirth;
    }

    public function getFirstName(): string {
        return $this->firstName;
    }

    public function getLastName(): string {
        return $this->lastName;
    }

    public function getFullName(): string {
        return $this->firstName . ' ' . $this->lastName;
    }

    public function getEmail(): string {
        return $this->email;
    }

    public function getPhone(): ?string {
        return $this->phone;
    }

    public function getDateOfBirth(): ?string {
        return $this->dateOfBirth;
    }

    public function isValid(): bool {
        if (empty($this->firstName) || empty($this->lastName)) {
            return false;
        }
        if (!filter_var($this->email, FILTER_VALIDATE_EMAIL)) {
            return false;
        }
        return true;
    }
}
