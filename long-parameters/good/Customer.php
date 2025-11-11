<?php

class Customer {
    public int $id;
    public string $name;
    public string $email;
    public ?string $phone;
    public Address $shippingAddress;
    public Address $billingAddress;

    public function __construct(
        int $id,
        string $name,
        string $email,
        ?string $phone = null,
        ?Address $shippingAddress = null,
        ?Address $billingAddress = null
    ) {
        $this->id = $id;
        $this->name = $name;
        $this->email = $email;
        $this->phone = $phone;
        $this->shippingAddress = $shippingAddress ?? new Address();
        $this->billingAddress = $billingAddress ?? new Address();
    }
}

class Address {
    public string $street = '';
    public string $city = '';
    public string $state = '';
    public string $zipCode = '';

    public function __construct(
        string $street = '',
        string $city = '',
        string $state = '',
        string $zipCode = ''
    ) {
        $this->street = $street;
        $this->city = $city;
        $this->state = $state;
        $this->zipCode = $zipCode;
    }
}

class Product {
    public int $id;
    public string $name;
    public float $price;

    public function __construct(int $id, string $name, float $price) {
        $this->id = $id;
        $this->name = $name;
        $this->price = $price;
    }
}

class OrderDetails {
    public Product $product;
    public int $quantity;
    public float $taxRate;
    public float $discountPercent;
    public string $shippingMethod;
    public float $shippingCost;
    public string $paymentMethod;
    public ?string $notes;

    public function __construct(
        Product $product,
        int $quantity,
        float $taxRate = 0.0,
        float $discountPercent = 0.0,
        string $shippingMethod = 'standard',
        float $shippingCost = 0.0,
        string $paymentMethod = 'credit_card',
        ?string $notes = null
    ) {
        $this->product = $product;
        $this->quantity = $quantity;
        $this->taxRate = $taxRate;
        $this->discountPercent = $discountPercent;
        $this->shippingMethod = $shippingMethod;
        $this->shippingCost = $shippingCost;
        $this->paymentMethod = $paymentMethod;
        $this->notes = $notes;
    }
}
